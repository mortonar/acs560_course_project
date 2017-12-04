using BookTracker.HelperClasses;
using System.Windows.Controls;
using System.Security;

namespace BookTracker
{
    /// <summary>
    /// Interaction logic for CreateAccountView.xaml
    /// </summary>
    public partial class CreateAccountView : UserControl, IPasswordProvider
    {
        public CreateAccountView()
        {
            InitializeComponent();
        }

        public SecureString Password
        {
            get { return UserPassword.SecurePassword;  }
        }
    }
}
