import socket,threading,time

allconn = {}

def handler(sock, addr):
	print('Accept new connection from %s:%s...' % addr)
	sock.send(b'Welcome!')
	while True:
		data = sock.recv(1024)
		time.sleep(1)
		if not data or data.decode('utf-8') == 'exit':
			break

		for oths in allconn:
			if oths != addr:
				allconn[oths].send(data)
	sock.close()
	print('Connection from %s:%s closed.' % addr)

s = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
s.bind(('127.0.0.1', 9999))
s.listen(5)

while True:
	sock, addr = s.accept()
	allconn[addr] = sock
	t = threading.Thread(target=handler, args=(sock, addr))
	t.start()

