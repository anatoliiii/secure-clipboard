# Shared Clipboard with LDAP Authentication Support

This project provides an implementation of a shared clipboard with LDAP authentication support. It consists of a client-side and a server-side component, enabling users to upload files to the server and manage clipboard contents.
## Description with Smiley Faces
Secure Clipboard Server is a Python-based application that provides a secure and convenient way to share clipboard content across multiple devices in a local network. It allows users to upload files, retrieve clipboard content, and set clipboard content remotely through a secure server.

Key Features:

- Secure communication: The server uses HTTPS to ensure secure communication and data transfer. :lock:
- LDAP Authentication: Users can authenticate themselves using LDAP credentials to ensure secure access to their clipboard. :closed_lock_with_key:
- File Uploads: Users can upload files to the server, making it easy to share files across devices. :file_folder:
- Clipboard Syncing: Users can retrieve and set clipboard content remotely, enabling seamless sharing of text across devices. :clipboard:
- Cross-platform Compatibility: The server can be run on various operating systems, including Linux, Windows, and macOS. :desktop_computer:

With Secure Clipboard Server, you can rest assured that your clipboard content is secure and conveniently accessible on any device within your local network. Enjoy the convenience of clipboard synchronization while maintaining the confidentiality and security of your data. :blush:✨

Start using Secure Clipboard Server today and experience a convenient and secure solution for clipboard data exchange.:+1::lock::clipboard::computer:


## Project Structure:

client: Python-based client-side component for interacting with the server, handling file uploads, and clipboard management.
server: Go-based server-side component responsible for handling client requests, authentication, file upload, and clipboard management.
ldap: Configuration file example for connecting to the LDAP server.
## Requirements:

Python 3.x
Python libraries: Flask, requests, clipboard, ldap3
Go 1.21+
## Installation and Setup:

1. Clone the repository: 
``` 
git clone https://github.com/your-username/clipboard-app.git
```
2. Navigate to the project directory: 
```
cd clipboard-app
```
3. Install the dependencies: 
```
pip install -r requirements.txt
```
4. Configure the LDAP server connection in the ```ldap/ldap_config.py``` file.
5. Open the necessary ports in your firewall for the LDAP server (usually port 389) and the server-side component (default port 5000).
6. Set the server's IP address in the local network as the SERVER_IP variable in the client/app.py file.
7. Start the server:
```
SECURE_CLIPBOARD_USERS="demo=demo=Demo User" go run ./cmd/server
```
8. Run the client application:
```
python client/app.py
```
## Usage:

- Uploading a file to the server:
  - Open the client application.
  - Enter your LDAP credentials (username and password).
  - Select a file to upload.
  - Click the "Upload" button.
  - Retrieving clipboard contents:

- Open the client application.
  - Enter your LDAP credentials.
  - Click the "Get Clipboard" button.
  - Setting clipboard contents:

* Open the client application.
  * Enter your LDAP credentials.
  * Enter the text to set in the clipboard.
  * Click the "Set Clipboard" button.
## Important:

- Ensure that the necessary ports for the LDAP server and the server-side component are open in your firewall to allow communication between the client and the server.
- Set the server's IP address in the local network as the SERVER_IP variable in the client/app.py file to ensure proper connection from the client to the server.
- Ensure the security of your infrastructure, including authentication, encryption, and protection against potential vulnerabilities.
## License
-------
This project is licensed under the [GNU Affero General Public License v3.0 (AGPL-3.0)]([LICENSE.md](https://github.com/anatoliiii/secure-clipboard-server/blob/main/LICENSE)). You can find the full text of the license [here]([LICENSE.md](https://github.com/anatoliiii/secure-clipboard-server/blob/main/LICENSE)).
## Copyright

© 2023 MOZGOLIKA

Email: mozgolika@vk.com
