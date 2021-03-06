//  Copyright (c) 2014 Couchbase, Inc.
//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file
//  except in compliance with the License. You may obtain a copy of the License at
//    http://www.apache.org/licenses/LICENSE-2.0
//  Unless required by applicable law or agreed to in writing, software distributed under the
//  License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
//  either express or implied. See the License for the specific language governing permissions
//  and limitations under the License.

package indexer

import (
	"errors"
	"fmt"
	"github.com/couchbase/indexing/secondary/common"
	"github.com/couchbase/indexing/secondary/dcp"
	"github.com/couchbase/indexing/secondary/logging"
	"net"
	"strconv"
	"time"
)

func IsIPLocal(ip string) bool {

	netIP := net.ParseIP(ip)

	//if loopback address, return true
	if netIP.IsLoopback() {
		return true
	}

	//compare with the local ip
	if localIP, err := GetLocalIP(); err == nil {
		if localIP.Equal(netIP) {
			return true
		}
	}

	return false

}

func GetLocalIP() (net.IP, error) {

	tt, err := net.Interfaces()
	if err != nil {
		return nil, err
	}
	for _, t := range tt {
		aa, err := t.Addrs()
		if err != nil {
			return nil, err
		}
		for _, a := range aa {
			ipnet, ok := a.(*net.IPNet)
			if !ok {
				continue
			}
			v4 := ipnet.IP.To4()
			if v4 == nil || v4[0] == 127 { // loopback address
				continue
			}
			return v4, nil
		}
	}
	return nil, errors.New("cannot find local IP address")
}

func IndexPath(inst *common.IndexInst, sliceId SliceId) string {
	return fmt.Sprintf("%s_%s_%d_%d.index", inst.Defn.Bucket, inst.Defn.Name, inst.InstId, sliceId)
}

func GetCurrentKVTs(bucket *couchbase.Bucket, numVbs int) (Timestamp, error) {
	ts := NewTimestamp(numVbs)
	start := time.Now()

	//get all the vb seqnum
	stats, err := bucket.GetStats("vbucket-details")
	//for all nodes in cluster
	for _, nodestat := range stats {
		//for all vbuckets
		for i := 0; i < numVbs; i++ {
			vbStateKey := "vb_" + strconv.Itoa(i)
			if state, ok := nodestat[vbStateKey]; ok {
				//only active vbuckets
				if state == "active" {
					vbSeqKey := "vb_" + strconv.Itoa(i) + ":high_seqno"
					if highseqno, ok := nodestat[vbSeqKey]; ok {
						if s, err := strconv.Atoi(highseqno); err == nil {
							ts[i] = Seqno(s)
						}
					}
				}
			}
		}
	}
	elapsed := time.Since(start)
	logging.Tracef("Indexer::getCurrentKVTs Time Taken %v \n\t TS Returned %v", elapsed, ts)
	if err != nil {
		logging.Errorf("Indexer::getCurrentKVTs Error Connecting to KV Cluster %v", err)
	}
	return ts, err
}

func ValidateBucket(cluster, bucket string) bool {

	var cinfo *common.ClusterInfoCache
	url, err := common.ClusterAuthUrl(cluster)
	if err == nil {
		cinfo, err = common.NewClusterInfoCache(url, DEFAULT_POOL)
	}
	if err != nil {
		logging.Fatalf("Indexer::Fail to init ClusterInfoCache : %v", err)
		common.CrashOnError(err)
	}

	cinfo.Lock()
	defer cinfo.Unlock()

	if err := cinfo.Fetch(); err != nil {
		logging.Errorf("Indexer::Fail to init ClusterInfoCache : %v", err)
		common.CrashOnError(err)
	}

	if nids, err := cinfo.GetNodesByBucket(bucket); err == nil && len(nids) != 0 {
		return true
	} else {
		return false
	}
}
