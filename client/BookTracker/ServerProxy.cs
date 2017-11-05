using BookTracker.Messaging.Requests;
using Newtonsoft.Json;
using System.Net.Sockets;
using System.Text;

namespace BookTracker
{
    public class ServerProxy
    {
        private static ServerProxy instance = null;
        private static readonly object mutex = new object();

        TcpClient clientSocket = new TcpClient();
        // TODO - implement a domain instead of hardcoded IP address to allow for changes in the VM's address
        public string ipAddress = "192.168.1.107";
        public int port = 8000;

        public static ServerProxy Instance
        {
            get
            {
                lock (mutex)
                {
                    if (instance == null)
                    {
                        instance = new ServerProxy();
                    }
                    return instance;
                }
            }
        }

        private ServerProxy()
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

        public string sendRequest(Base message)
        {
            NetworkStream serverStream = clientSocket.GetStream();

            string request = JsonConvert.SerializeObject(message);

            byte[] outStream = Encoding.ASCII.GetBytes(request);
            serverStream.Write(outStream, 0, outStream.Length);
            serverStream.Flush();

            // TODO - This needs fixed and cleaned up once the server is returning data
            byte[] inStream = new byte[(int)clientSocket.ReceiveBufferSize];
            serverStream.Read(inStream, 0, (int)clientSocket.ReceiveBufferSize);
            string returndata = Encoding.ASCII.GetString(inStream);
            return returndata;
        }
    }

}

