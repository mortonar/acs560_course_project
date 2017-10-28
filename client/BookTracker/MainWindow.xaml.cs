using System.Windows;

namespace BookTracker
{
    /// <summary>
    /// Interaction logic for MainWindow.xaml
    /// </summary>
    public partial class MainWindow : Window
    {
        public ServerProxy serverProxy = new ServerProxy();

        public MainWindow()
        {
            InitializeComponent();
        }

        // Button for testing server requests
        private void Button_Click(object sender, RoutedEventArgs e)
        {
            serverProxy.sendRequest();
        }
    }
}
