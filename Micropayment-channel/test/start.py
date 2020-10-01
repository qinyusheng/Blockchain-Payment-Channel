#!/usr/bin/python

from mininet.net import Mininet
from mininet.node import CPULimitedHost
from mininet.cli import CLI
from mininet.topo import Topo
from mininet.link import TCLink
from mininet.log import info
from mininet.log import setLogLevel

cpuPercent = 0.5
bandwidth = 20
delay = '100ms'

class P2PTopo(Topo):
    def build(self):
    	switch = self.addSwitch('s0')
        
    	ipfs = self.addHost('ipfs')
    	self.addLink(ipfs, switch, bw=bandwidth, delay=delay)
        
    	for i in range(2):
    		host = self.addHost('h%s' % i, cpu=cpuPercent)
    		self.addLink(host, switch, bw=bandwidth, delay=delay)

def runNet():
    topo = P2PTopo()
    net = Mininet(topo=topo, host=CPULimitedHost, link=TCLink)
    net.start()
    startServer(net)
    CLI(net)
    stopServer(net)
    net.stop()

def startServer(net):
    ipfs = net.get('ipfs')
    host0 = net.get('h0')
    host1 = net.get('h1')
    
    info('*** IPFS node starting on %s\n' % ipfs)
    info('./ipfs\n')
#    ipfs.cmd('./ipfs &')
    
    info('*** Blockchain nodes starting on %s\n' % host0)
    info('./nodes --ipfs %s \n' % ipfs.IP())
#    host0.cmd('./nodes -i', host0.IP(), '-p 9000 --peer %s' % host1.IP(), '--ipfs %s &' % ipfs.IP())
    
    info('*** Blockchain nodec starting on %s\n' % host1)
    info('./nodec --peer %s' % host0.IP(), '--ipfs %s\n' % ipfs.IP())
#    host1.cmd('./nodec -i', host1.IP(), '-p 9000 --peer %s' % host0.IP(), '--ipfs %s' % ipfs.IP())

def stopServer(net):
    for h in net.hosts:
        info('*** Node stopping on %s\n' % h)

if __name__ == '__main__':
    setLogLevel('info')
    runNet()
    info( 'Done.\n')
