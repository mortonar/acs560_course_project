using BookTracker.HelperClasses;
using BookTracker.Messaging.Requests;
using BookTracker.Models;
using Newtonsoft.Json;
using Newtonsoft.Json.Linq;
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
            Debug.WriteLine("Logging in with credentials: " + _loginModel.UserName + " and " + LoginModel.Password);
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
            try
            {
                string response = ServerProxy.Instance.sendRequest(message);
                Debug.WriteLine("RESPONSE: " + response);

                Messaging.Responses.Base resp = JsonConvert.DeserializeObject<Messaging.Responses.Base>(response);
                if (resp.Success)
                {
                    Messaging.Responses.Login loginResp = (resp.Payload as JObject).ToObject<Messaging.Responses.Login>();
                    Debug.WriteLine("Setting token: " + resp.Success + " | " + loginResp.Token);
                    ((App)App.Current).setToken(loginResp.Token);
                    Debug.WriteLine("Logged in? " + ((App)App.Current).isLoggedIn());
                    ((App)App.Current).changeViewModel(new ReadingListViewModel());
                }
            }
            catch (Exception e)
            {
                Debug.WriteLine("Exception: " + e);
            }
        }

        public void Create_Account()
        {
            ((App)App.Current).changeViewModel(new CreateAccountViewModel());
        }

        public bool CanExecute(object parameter)
        {
            return true;
        }

        public void Execute(object parameter)
        {
            if (parameter.GetType() == typeof(String) && parameter.Equals("Create")) {
                Create_Account();
            } else {
                IPasswordProvider passwordProv = parameter as IPasswordProvider;
                LoginModel.Password = PasswordUtils.ConvertToUnsecureString(passwordProv.Password);
                Login();
            }
        }

        public void Update() { }

    }
}