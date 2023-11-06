docker run -itd --restart=unless-stopped -p 2379:2379 --net etcd --ip 172.28.5.1 \
  -p 2380:2380 --name etcd0 quay.io/coreos/etcd:latest \
  /usr/local/bin/etcd --data-dir=/etcd-data --name etcd-node-0 \
  --initial-advertise-peer-urls http://172.28.5.1:2380 \
  --listen-peer-urls http://0.0.0.0:2380 \
  --advertise-client-urls http://172.28.5.1:2379 \
  --listen-client-urls http://0.0.0.0:2379 \
  --initial-cluster etcd-node-0=http://172.28.5.1:2380,etcd-node-1=http://172.28.5.2:2381,etcd-node-2=http://172.28.5.3:2382 \
  --initial-cluster-state new --initial-cluster-token my-etcd-token

docker run -itd --restart=unless-stopped -p 2340:2340 --net etcd --ip 172.28.5.2 \
  -p 2381:2381 --name etcd1 quay.io/coreos/etcd:latest \
  /usr/local/bin/etcd --data-dir=/etcd-data --name etcd-node-1 \
  --initial-advertise-peer-urls http://172.28.5.2:2381 \
  --listen-peer-urls http://0.0.0.0:2381 \
  --advertise-client-urls http://172.28.5.2:2340 \
  --listen-client-urls http://0.0.0.0:2340 \
  --initial-cluster etcd-node-0=http://172.28.5.1:2380,etcd-node-1=http://172.28.5.2:2381,etcd-node-2=http://172.28.5.3:2382 \
  --initial-cluster-state new --initial-cluster-token my-etcd-token

docker run -itd --restart=unless-stopped -p 2341:2341 --net etcd --ip 172.28.5.3 \
  -p 2382:2382 --name etcd2 quay.io/coreos/etcd:latest \
  /usr/local/bin/etcd --data-dir=/etcd-data --name etcd-node-2 \
  --initial-advertise-peer-urls http://172.28.5.3:2382 \
  --listen-peer-urls http://0.0.0.0:2382 \
  --advertise-client-urls http://172.28.5.3:2341 \
  --listen-client-urls http://0.0.0.0:2341 \
  --initial-cluster etcd-node-0=http://172.28.5.1:2380,etcd-node-1=http://172.28.5.2:2381,etcd-node-2=http://172.28.5.3:2382 \
  --initial-cluster-state new --initial-cluster-token my-etcd-token

docker ps -a

etcdctl member list -w table

etcdctl --endpoints 127.0.0.1:2340,127.0.0.1:2341,127.0.0.1:2379 endpoint status -w table
export ETCDCTL_ENDPOINTS="http://172.28.5.1:2379,http://172.28.5.2:2379,http://172.28.5.3:2379"
#etcdctl get --keys-only --prefix "" for finding all keys
