# A Python program to show server/client model by using Go as server and Python as the client side.
# This program should run after 'go_server.go' Go program starts, to test the interoperability.

import socket

def main():
    server_ip = input("Enter server IP: ")

    client = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
    client.connect((server_ip, 8080))

    while True:
        message = input("Enter message: ")
        client.sendall(message.encode())

        data = client.recv(1024)
        print(f"Received response: {data.decode()}")

if __name__ == "__main__":
    main()