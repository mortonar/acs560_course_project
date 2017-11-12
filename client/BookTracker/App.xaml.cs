using System.Diagnostics;
using System.Windows;

namespace BookTracker
{
    /// <summary>
    /// Interaction logic for App.xaml
    /// </summary>
    public partial class App : Application
    {

        private MainWindowViewModel context;

        void App_Startup(object sender, StartupEventArgs sea)
        {
            Debug.WriteLine("Starting up...");

            MainWindow app = new MainWindow();
            context = new MainWindowViewModel();
            app.DataContext = context;
            app.Show();

            // TODO use command line arg to set ServerProxy's server address
            // initialize the ServerProxy
            ServerProxy sp = ServerProxy.Instance;
        }

        public void changeViewModel(IPageViewModel viewModel)
        {
            context.ChangePageCommand.Execute(viewModel);
        }

    }

}
