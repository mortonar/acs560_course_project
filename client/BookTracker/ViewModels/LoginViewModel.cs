using BookTracker.HelperClasses;
using BookTracker.Messaging.Requests;
using BookTracker.Models;
using System;
using System.Diagnostics;
using System.Windows.Input;

namespace BookTracker.ViewModels
{
    public class LoginViewModel : ObservableObject, ICommand, IPageViewModel
    {

        private LoginModel _loginModel;

        public event EventHandler CanExecuteChanged;

        public LoginModel LoginModel
        {
            get { return _loginModel; }
            set { _loginModel = value; }
        }

        public string Name
        {
            get { return "Login"; }
        }

        public LoginViewModel()
        {
            _loginModel = new LoginModel();
        }

        public void Login()
        {
            Debug.Write("Logging in with credentials: " + _loginModel.UserName + " and " + LoginModel.Password);
            Login login = new Login
            {
                UserName = _loginModel.UserName,
                EncryptedPass = _loginModel.Password
            };
            Base message = new Base
            {
                Action = "Auth",
                Payload = login
            };
            string response = ServerProxy.Instance.sendRequest(message);
            Debug.WriteLine("RESPONSE: " + response);
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