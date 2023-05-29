sudo ip netns add router1
sudo ip netns add host1
sudo ip netns add host2

sudo ip link add host1-router1 type veth peer name router1-host1
sudo ip link add host2-router1 type veth peer name router1-host2

sudo ip link set host1-router1 netns host1
sudo ip link set router1-host1 netns router1
sudo ip link set host2-router1 netns host2
sudo ip link set router1-host2 netns router1

sudo ip netns exec host1 ip link set host1-router1 up
sudo ip netns exec router1 ip link set router1-host1 up
sudo ip netns exec host2 ip link set host2-router1 up
sudo ip netns exec router1 ip link set router1-host2 up

sudo ip netns exec host1 ip addr add 10.0.1.2/24 dev host1-router1
sudo ip netns exec router1 ip addr add 10.0.1.254/24 dev router1-host1
sudo ip netns exec host2 ip addr add 192.0.1.2/24 dev host2-router1
sudo ip netns exec router1 ip addr add 192.0.1.254/24 dev router1-host2

sudo ip netns exec host1 ip route add 192.0.1.2 via 10.0.1.254
sudo ip netns exec host2 ip route add 10.0.1.2 via 192.0.1.254

sudo ip netns exec router1 sysctl net.ipv4.ip_forward=1

# お掃除
# sudo ip netns exec host1 ip link del host1-router1;sudo ip netns exec host2 ip link del host2-router1;sudo ip netns del host1;sudo ip netns del host2;sudo ip netns del router1

# 疎通確認
# sudo ip netns exec host1 ping -c 2 192.0.1.2

# arp確認 c=回数 
# sudo ip netns exec host1 arping -c 1 -I host1-router1 10.0.1.254

# ルーター起動
# sudo ip netns exec router1 ./main --mode ch1