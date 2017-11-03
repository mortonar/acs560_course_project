using System;
using System.Diagnostics;
using System.Windows.Input;

namespace BookTracker
{
    public class LoginViewModel : ICommand
    {

        private LoginModel _loginModel;

        public event EventHandler CanExecuteChanged;

        public LoginModel LoginModel
        {
            get { return _loginModel; }
            set {
                _loginModel = value;
            }
        }

        public LoginViewModel()
        {
            _loginModel = new LoginModel();
        }

        public void Login()
        {
            Debug.Write("Logging in with credentials: " + _loginModel.UserName + " and " + LoginModel.Password);
        }

        public bool CanExecute(object parameter)
        {
            return true;
        }

        public void Execute(object parameter)
        {
            Login();
        }
    }
}