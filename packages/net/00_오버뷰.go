/*
http://golang.org/pkg/net/

net패키지는 네트워크 입출력을 위한 휴대용(portable) 인터페이스를 제공합니다.
TCP/IP , UDP , 도메인네임 추출(원래 영어단어는 resolution으로 적었다), 유닉스 도메인소켓같은 것들을 포함합니다.

비록 패키지가 low레벨 네트워크 요소들에 대한 접근을 제공하지만,
대부분의 클라이언트들은 Dial, Listen, Accect function들이 제공하는
기초적 인터페이스들만을 필요로 할 것입니다. (맞게 적었나;;)


crypto/tls 패키지는 같은 인터페이스를 사용하고, 비슷한 Dial , Listen function을 이용합니다.
Dial function은 서버로 접속합니다.

*/
conn, err := net.Dial("tcp", "google.com:80")
if err != nil {
	// handle error
}
fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")
status, err := bufio.NewReader(conn).ReadString('\n')
// ...

/* Listen function은 서버를 생성합니다 */
ln, err := net.Listen("tcp", ":8080")
if err != nil {
	// handle error
}
for {
	conn, err := ln.Accept()
	if err != nil {
		// handle error
	}
	go handleConnection(conn)
}

/*
   API

      Constants
      Variables
      func InterfaceAddrs() ([]Addr, error)
      func Interfaces() ([]Interface, error)
      func JoinHostPort(host, port string) string
      func LookupAddr(addr string) (name []string, err error)
      func LookupCNAME(name string) (cname string, err error)
      func LookupHost(host string) (addrs []string, err error)
      func LookupIP(host string) (addrs []IP, err error)
      func LookupMX(name string) (mx []*MX, err error)
      func LookupNS(name string) (ns []*NS, err error)
      func LookupPort(network, service string) (port int, err error)
      func LookupSRV(service, proto, name string) (cname string, addrs []*SRV, err error)
      func LookupTXT(name string) (txt []string, err error)
      func SplitHostPort(hostport string) (host, port string, err error)
      type Addr
      type AddrError
          func (e *AddrError) Error() string
          func (e *AddrError) Temporary() bool
          func (e *AddrError) Timeout() bool
      type Conn
          func Dial(network, address string) (Conn, error)
          func DialTimeout(network, address string, timeout time.Duration) (Conn, error)
          func FileConn(f *os.File) (c Conn, err error)
          func Pipe() (Conn, Conn)
      type DNSConfigError
          func (e *DNSConfigError) Error() string
          func (e *DNSConfigError) Temporary() bool
          func (e *DNSConfigError) Timeout() bool
      type DNSError
          func (e *DNSError) Error() string
          func (e *DNSError) Temporary() bool
          func (e *DNSError) Timeout() bool
      type Dialer
          func (d *Dialer) Dial(network, address string) (Conn, error)
      type Error
      type Flags
          func (f Flags) String() string
      type HardwareAddr
          func ParseMAC(s string) (hw HardwareAddr, err error)
          func (a HardwareAddr) String() string
      type IP
          func IPv4(a, b, c, d byte) IP
          func ParseCIDR(s string) (IP, *IPNet, error)
          func ParseIP(s string) IP
          func (ip IP) DefaultMask() IPMask
          func (ip IP) Equal(x IP) bool
          func (ip IP) IsGlobalUnicast() bool
          func (ip IP) IsInterfaceLocalMulticast() bool
          func (ip IP) IsLinkLocalMulticast() bool
          func (ip IP) IsLinkLocalUnicast() bool
          func (ip IP) IsLoopback() bool
          func (ip IP) IsMulticast() bool
          func (ip IP) IsUnspecified() bool
          func (ip IP) MarshalText() ([]byte, error)
          func (ip IP) Mask(mask IPMask) IP
          func (ip IP) String() string
          func (ip IP) To16() IP
          func (ip IP) To4() IP
          func (ip *IP) UnmarshalText(text []byte) error
      type IPAddr
          func ResolveIPAddr(net, addr string) (*IPAddr, error)
          func (a *IPAddr) Network() string
          func (a *IPAddr) String() string
      type IPConn
          func DialIP(netProto string, laddr, raddr *IPAddr) (*IPConn, error)
          func ListenIP(netProto string, laddr *IPAddr) (*IPConn, error)
          func (c *IPConn) Close() error
          func (c *IPConn) File() (f *os.File, err error)
          func (c *IPConn) LocalAddr() Addr
          func (c *IPConn) Read(b []byte) (int, error)
          func (c *IPConn) ReadFrom(b []byte) (int, Addr, error)
          func (c *IPConn) ReadFromIP(b []byte) (int, *IPAddr, error)
          func (c *IPConn) ReadMsgIP(b, oob []byte) (n, oobn, flags int, addr *IPAddr, err error)
          func (c *IPConn) RemoteAddr() Addr
          func (c *IPConn) SetDeadline(t time.Time) error
          func (c *IPConn) SetReadBuffer(bytes int) error
          func (c *IPConn) SetReadDeadline(t time.Time) error
          func (c *IPConn) SetWriteBuffer(bytes int) error
          func (c *IPConn) SetWriteDeadline(t time.Time) error
          func (c *IPConn) Write(b []byte) (int, error)
          func (c *IPConn) WriteMsgIP(b, oob []byte, addr *IPAddr) (n, oobn int, err error)
          func (c *IPConn) WriteTo(b []byte, addr Addr) (int, error)
          func (c *IPConn) WriteToIP(b []byte, addr *IPAddr) (int, error)
      type IPMask
          func CIDRMask(ones, bits int) IPMask
          func IPv4Mask(a, b, c, d byte) IPMask
          func (m IPMask) Size() (ones, bits int)
          func (m IPMask) String() string
      type IPNet
          func (n *IPNet) Contains(ip IP) bool
          func (n *IPNet) Network() string
          func (n *IPNet) String() string
      type Interface
          func InterfaceByIndex(index int) (*Interface, error)
          func InterfaceByName(name string) (*Interface, error)
          func (ifi *Interface) Addrs() ([]Addr, error)
          func (ifi *Interface) MulticastAddrs() ([]Addr, error)
      type InvalidAddrError
          func (e InvalidAddrError) Error() string
          func (e InvalidAddrError) Temporary() bool
          func (e InvalidAddrError) Timeout() bool
      type Listener
          func FileListener(f *os.File) (l Listener, err error)
          func Listen(net, laddr string) (Listener, error)
      type MX
      type NS
      type OpError
          func (e *OpError) Error() string
          func (e *OpError) Temporary() bool
          func (e *OpError) Timeout() bool
      type PacketConn
          func FilePacketConn(f *os.File) (c PacketConn, err error)
          func ListenPacket(net, laddr string) (PacketConn, error)
      type ParseError
          func (e *ParseError) Error() string
      type SRV
      type TCPAddr
          func ResolveTCPAddr(net, addr string) (*TCPAddr, error)
          func (a *TCPAddr) Network() string
          func (a *TCPAddr) String() string
      type TCPConn
          func DialTCP(net string, laddr, raddr *TCPAddr) (*TCPConn, error)
          func (c *TCPConn) Close() error
          func (c *TCPConn) CloseRead() error
          func (c *TCPConn) CloseWrite() error
          func (c *TCPConn) File() (f *os.File, err error)
          func (c *TCPConn) LocalAddr() Addr
          func (c *TCPConn) Read(b []byte) (int, error)
          func (c *TCPConn) ReadFrom(r io.Reader) (int64, error)
          func (c *TCPConn) RemoteAddr() Addr
          func (c *TCPConn) SetDeadline(t time.Time) error
          func (c *TCPConn) SetKeepAlive(keepalive bool) error
          func (c *TCPConn) SetKeepAlivePeriod(d time.Duration) error
          func (c *TCPConn) SetLinger(sec int) error
          func (c *TCPConn) SetNoDelay(noDelay bool) error
          func (c *TCPConn) SetReadBuffer(bytes int) error
          func (c *TCPConn) SetReadDeadline(t time.Time) error
          func (c *TCPConn) SetWriteBuffer(bytes int) error
          func (c *TCPConn) SetWriteDeadline(t time.Time) error
          func (c *TCPConn) Write(b []byte) (int, error)
      type TCPListener
          func ListenTCP(net string, laddr *TCPAddr) (*TCPListener, error)
          func (l *TCPListener) Accept() (Conn, error)
          func (l *TCPListener) AcceptTCP() (*TCPConn, error)
          func (l *TCPListener) Addr() Addr
          func (l *TCPListener) Close() error
          func (l *TCPListener) File() (f *os.File, err error)
          func (l *TCPListener) SetDeadline(t time.Time) error
      type UDPAddr
          func ResolveUDPAddr(net, addr string) (*UDPAddr, error)
          func (a *UDPAddr) Network() string
          func (a *UDPAddr) String() string
      type UDPConn
          func DialUDP(net string, laddr, raddr *UDPAddr) (*UDPConn, error)
          func ListenMulticastUDP(net string, ifi *Interface, gaddr *UDPAddr) (*UDPConn, error)
          func ListenUDP(net string, laddr *UDPAddr) (*UDPConn, error)
          func (c *UDPConn) Close() error
          func (c *UDPConn) File() (f *os.File, err error)
          func (c *UDPConn) LocalAddr() Addr
          func (c *UDPConn) Read(b []byte) (int, error)
          func (c *UDPConn) ReadFrom(b []byte) (int, Addr, error)
          func (c *UDPConn) ReadFromUDP(b []byte) (n int, addr *UDPAddr, err error)
          func (c *UDPConn) ReadMsgUDP(b, oob []byte) (n, oobn, flags int, addr *UDPAddr, err error)
          func (c *UDPConn) RemoteAddr() Addr
          func (c *UDPConn) SetDeadline(t time.Time) error
          func (c *UDPConn) SetReadBuffer(bytes int) error
          func (c *UDPConn) SetReadDeadline(t time.Time) error
          func (c *UDPConn) SetWriteBuffer(bytes int) error
          func (c *UDPConn) SetWriteDeadline(t time.Time) error
          func (c *UDPConn) Write(b []byte) (int, error)
          func (c *UDPConn) WriteMsgUDP(b, oob []byte, addr *UDPAddr) (n, oobn int, err error)
          func (c *UDPConn) WriteTo(b []byte, addr Addr) (int, error)
          func (c *UDPConn) WriteToUDP(b []byte, addr *UDPAddr) (int, error)
      type UnixAddr
          func ResolveUnixAddr(net, addr string) (*UnixAddr, error)
          func (a *UnixAddr) Network() string
          func (a *UnixAddr) String() string
      type UnixConn
          func DialUnix(net string, laddr, raddr *UnixAddr) (*UnixConn, error)
          func ListenUnixgram(net string, laddr *UnixAddr) (*UnixConn, error)
          func (c *UnixConn) Close() error
          func (c *UnixConn) CloseRead() error
          func (c *UnixConn) CloseWrite() error
          func (c *UnixConn) File() (f *os.File, err error)
          func (c *UnixConn) LocalAddr() Addr
          func (c *UnixConn) Read(b []byte) (int, error)
          func (c *UnixConn) ReadFrom(b []byte) (int, Addr, error)
          func (c *UnixConn) ReadFromUnix(b []byte) (n int, addr *UnixAddr, err error)
          func (c *UnixConn) ReadMsgUnix(b, oob []byte) (n, oobn, flags int, addr *UnixAddr, err error)
          func (c *UnixConn) RemoteAddr() Addr
          func (c *UnixConn) SetDeadline(t time.Time) error
          func (c *UnixConn) SetReadBuffer(bytes int) error
          func (c *UnixConn) SetReadDeadline(t time.Time) error
          func (c *UnixConn) SetWriteBuffer(bytes int) error
          func (c *UnixConn) SetWriteDeadline(t time.Time) error
          func (c *UnixConn) Write(b []byte) (int, error)
          func (c *UnixConn) WriteMsgUnix(b, oob []byte, addr *UnixAddr) (n, oobn int, err error)
          func (c *UnixConn) WriteTo(b []byte, addr Addr) (n int, err error)
          func (c *UnixConn) WriteToUnix(b []byte, addr *UnixAddr) (n int, err error)
      type UnixListener
          func ListenUnix(net string, laddr *UnixAddr) (*UnixListener, error)
          func (l *UnixListener) Accept() (c Conn, err error)
          func (l *UnixListener) AcceptUnix() (*UnixConn, error)
          func (l *UnixListener) Addr() Addr
          func (l *UnixListener) Close() error
          func (l *UnixListener) File() (f *os.File, err error)
          func (l *UnixListener) SetDeadline(t time.Time) (err error)
      type UnknownNetworkError
          func (e UnknownNetworkError) Error() string
          func (e UnknownNetworkError) Temporary() bool
          func (e UnknownNetworkError) Timeout() bool
      Bugs

*/