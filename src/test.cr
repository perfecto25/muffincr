require "socket"

MULTICAST_IP = "224.0.25.121"
PORT = 21910
INTERFACE_IP = "192.168.38.38"


udp = UDPSocket.new(Socket::Family::INET)
udp.bind(INTERFACE_IP, 22499)
addr = Socket::IPAddress.new(MULTICAST_IP, PORT)
udp.join_group(addr)
#spawn do
  loop do
    puts udp.receive
    message, client_addr = udp.receive
    t=Time.utc
    puts "#{Time.utc(t.year,t.month,t.day,t.hour,t.minute,t.second).to_rfc3339} #{client_addr.address} #{message}"
  end
#end
Signal::INT.trap do
  udp.leave_group(addr)
  udp.close
end
