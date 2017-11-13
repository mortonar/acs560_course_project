using BookTracker.HelperClasses;
using System.Windows.Controls;
using System.Security;

namespace BookTracker
{
    /// <summary>
    /// Interaction logic for LoginView.xaml
    /// </summary>
    public partial class LoginView : UserControl, IPasswordProvider
    {
        public LoginView()
        {
            InitializeComponent();
        }

        public SecureString Password
        {
            get { return UserPassword.SecurePassword;  }
        }
    }
}
