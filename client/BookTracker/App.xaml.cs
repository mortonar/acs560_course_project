using System;
using System.Collections.Generic;
using System.Configuration;
using System.Data;
using System.Diagnostics;
using System.Linq;
using System.Threading.Tasks;
using System.Windows;

namespace BookTracker
{
    /// <summary>
    /// Interaction logic for App.xaml
    /// </summary>
    public partial class App : Application
    {
        void App_Startup(object sender, StartupEventArgs sea)
        {
            Debug.WriteLine("Starting up...");

            // TODO use command line arg to set ServerProxy's server address
            // initialize the ServerProxy
            ServerProxy sp = ServerProxy.Instance;
        }

    }

}
