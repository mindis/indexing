// Copyright (c) 2014 Couchbase, Inc.
// Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file
// except in compliance with the License. You may obtain a copy of the License at
//   http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software distributed under the
// License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
// either express or implied. See the License for the specific language governing permissions
// and limitations under the License.

package main

import (
	"flag"
	"os"
	"strings"

	"github.com/couchbase/cbauth"
	"github.com/couchbase/indexing/secondary/common"
	"github.com/couchbase/indexing/secondary/fdb"
	"github.com/couchbase/indexing/secondary/indexer"
	"github.com/couchbase/indexing/secondary/logging"
)

var (
	logLevel          = flag.String("loglevel", "Info", "Log Level - Silent, Fatal, Error, Info, Debug, Trace")
	numVbuckets       = flag.Int("vbuckets", indexer.MAX_NUM_VBUCKETS, "Number of vbuckets configured in Couchbase")
	cluster           = flag.String("cluster", indexer.DEFAULT_CLUSTER_ENDPOINT, "Couchbase cluster address")
	adminPort         = flag.String("adminPort", "9100", "Index ddl and status port")
	scanPort          = flag.String("scanPort", "9101", "Index scanner port")
	httpPort          = flag.String("httpPort", "9102", "Index http mgmt port")
	streamInitPort    = flag.String("streamInitPort", "9103", "Index initial stream port")
	streamCatchupPort = flag.String("streamCatchupPort", "9104", "Index catchup stream port")
	streamMaintPort   = flag.String("streamMaintPort", "9105", "Index maintenance stream port")
	storageDir        = flag.String("storageDir", "./", "Index file storage directory path")
	enableManager     = flag.Bool("enable_manager", true, "Enable Index Manager")
	auth              = flag.String("auth", "", "Auth user and password")

	// so we don't need to sync with ns_server for merge. remove this soon
	unused = flag.String("log", "", "Ignored")
)

func main() {
	common.HideConsole(true)
	defer common.HideConsole(false)
	common.SeedProcess()

	logging.Infof("Indexer started with command line: %v\n", os.Args)
	flag.Parse()

	logging.SetLogLevel(logging.Level(*logLevel))
	forestdb.Log = &logging.SystemLogger

	// setup cbauth
	if *auth != "" {
		up := strings.Split(*auth, ":")
		logging.Tracef("Initializing cbauth with user %v for cluster %v\n", up[0], *cluster)
		if _, err := cbauth.InternalRetryDefaultInit(*cluster, up[0], up[1]); err != nil {
			logging.Fatalf("Failed to initialize cbauth: %s", err)
		}
	}

	go common.DumpOnSignal()
	go common.ExitOnStdinClose()

	config := common.SystemConfig
	config.SetValue("indexer.clusterAddr", *cluster)
	config.SetValue("indexer.numVbuckets", *numVbuckets)
	config.SetValue("indexer.enableManager", *enableManager)
	config.SetValue("indexer.adminPort", *adminPort)
	config.SetValue("indexer.scanPort", *scanPort)
	config.SetValue("indexer.httpPort", *httpPort)
	config.SetValue("indexer.streamInitPort", *streamInitPort)
	config.SetValue("indexer.streamCatchupPort", *streamCatchupPort)
	config.SetValue("indexer.streamMaintPort", *streamMaintPort)
	config.SetValue("indexer.storage_dir", *storageDir)

	_, msg := indexer.NewIndexer(config)

	if msg.GetMsgType() != indexer.MSG_SUCCESS {
		logging.Warnf("Indexer Failure to Init %v", msg)
	}

	logging.Infof("Indexer exiting normally\n")
}
