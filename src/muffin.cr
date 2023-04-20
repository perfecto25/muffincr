require "socket"

struct Socket::IPAddress
  def addr
    @addr
  end
end

lib LibC
  struct IpvMreq
    ipvmr_multiaddr : LibC::InAddr
    ipvmr_interface : LibC::UInt
  end
end

s = UDPSocket.new Socket::Family::INET
s.bind "0.0.0.0", 0
# multicast_interface selection
# s.setsockopt 9, 1, 41
# s.getsockopt 9, 41

# 224.0.25.121:21910  tch4

sock = Socket::IPAddress.new("224.0.25.121", 21910)
addr = sock.addr.not_nil!

req = LibC::IpvMreq.new
req.ipvmr_multiaddr = addr
req.ipvmr_interface = 0
# Join IPv4 group

s.setsockopt 12, req, 41
s.send("testing", sock)
s.receive[0]
