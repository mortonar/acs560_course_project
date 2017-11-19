using System.Windows;

namespace BookTracker
{
    /// <summary>
    /// Interaction logic for MainWindow.xaml
    /// </summary>
    public partial class MainWindow : Window
    {
        public MainWindow()
        {
            InitializeComponent();
        }

        // Button for testing server requests
        private void Button_Click(object sender, RoutedEventArgs e)
        {
            ServerProxy.Instance.sendRequest(new Messaging.Requests.Base { });
        }

        private void Window_Unloaded(object sender, RoutedEventArgs e)
        {

        }
    }
}
