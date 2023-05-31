# Shared Clipboard with LDAP Authentication Support

This project provides an implementation of a shared clipboard with LDAP authentication support. It consists of a client-side and a server-side component, enabling users to upload files to the server and manage clipboard contents.

## Project Structure:

client: Python-based client-side component for interacting with the server, handling file uploads, and clipboard management.
server: Python-based server-side component using the Flask framework, responsible for handling client requests, authentication, file upload, and clipboard management.
ldap: Configuration file example for connecting to the LDAP server.
## Requirements:

Python 3.x
Python libraries: Flask, requests, clipboard, ldap3
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
python server/app.py
```
8. Run the client application: 
```
python client/app.py
```
## Usage:

1. Uploading a file to the server:

Open the client application.
Enter your LDAP credentials (username and password).
Select a file to upload.
Click the "Upload" button.
Retrieving clipboard contents:

2. Open the client application.
Enter your LDAP credentials.
Click the "Get Clipboard" button.
Setting clipboard contents:

3. Open the client application.
Enter your LDAP credentials.
Enter the text to set in the clipboard.
Click the "Set Clipboard" button.
## Important:

Ensure that the necessary ports for the LDAP server and the server-side component are open in your firewall to allow communication between the client and the server.
Set the server's IP address in the local network as the SERVER_IP variable in the client/app.py file to ensure proper connection from the client to the server.
Ensure the security of your infrastructure, including authentication, encryption, and protection against potential vulnerabilities.
## License:

This project is licensed under the GNU Affero General Public License v3.0 (AGPL-3.0).
