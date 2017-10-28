using System.Net.Sockets;
using System.Text;

namespace BookTracker
{
    public class ServerProxy
    {
        TcpClient clientSocket = new TcpClient();
        // TODO - implement a domain instead of hardcoded IP address to allow for changes in the VM's address
        public string ipAddress = "192.168.1.3";
        public int port = 8000;

        public ServerProxy()
        {
            try
            {
                clientSocket.Connect(ipAddress, port);
            }
            catch (SocketException SE)
            {
                string error = "An error occured while connecting [" + SE.Message + "]\n";
            }
        }

        public string sendRequest()
        {
            NetworkStream serverStream = clientSocket.GetStream();

            // TODO - implement request based off actual values
            string request = "{\"token\": \"asdf\",\"action\": \"Search\",\"payload\": {\"bookName\": \"Dying of the light\", \"authorName\": \"George R.R. Martin\"}}";

            byte[] outStream = Encoding.ASCII.GetBytes(request);
            serverStream.Write(outStream, 0, outStream.Length);
            serverStream.Flush();

            // TODO - This needs fixed and cleaned up once the server is returning data
            byte[] inStream = new byte[10025];
            serverStream.Read(inStream, 0, (int)clientSocket.ReceiveBufferSize);
            string returndata = Encoding.ASCII.GetString(inStream);
            return returndata;
        }
    }

}

