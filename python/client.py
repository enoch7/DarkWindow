import socket,select,threading,time

s = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
# 建立连接:
s.connect(('127.0.0.1', 9999))

def HandleInput(sock):
	pass

def HandleReceive(sock):
	while True:
		data = sock.recv(1024).decode('utf-8')
		print(data)
		time.sleep(1)


tInput = threading.Thread(target=HandleInput, args=(s, ))
tReceive = threading.Thread(target=HandleReceive, args=(s, ))

tInput.start()
tReceive.start()

tInput.join()
tReceive.join()