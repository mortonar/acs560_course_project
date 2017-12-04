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
    public class CreateAccountViewModel : ObservableObject, ICommand, IPageViewModel
    {

        private CreateAccountModel _createAccountModel;

        public event EventHandler CanExecuteChanged;

        public CreateAccountModel CreateAccountModel
        {
            get { return _createAccountModel; }
            set { _createAccountModel = value; }
        }

        public string Name
        {
            get { return "Create Account"; }
        }

        public CreateAccountViewModel()
        {
            _createAccountModel = new CreateAccountModel();
        }

        public void Create_Account()
        {
            Debug.Write("Create Account with credentials: " + _createAccountModel.UserName + " and " + CreateAccountModel.Password);
            CreateAccount createAccount = new CreateAccount
            {
                UserName = _createAccountModel.UserName,
                Email = _createAccountModel.Email,
                Password = _createAccountModel.Password
            };
            Base message = new Base
            {
                Action = "CreateAccount",
                Payload = createAccount
            };
            try
            {
                string rawResp = ServerProxy.Instance.sendRequest(message);
                Debug.WriteLine("RESPONSE: " + rawResp + "\n");
                ((App)App.Current).changeViewModel(new LoginViewModel());

            }
            catch (Exception e)
            {
                Debug.WriteLine("Exception: " + e);
            }
        }


        public bool CanExecute(object parameter)
        {
            return true;
        }

        public void Execute(object parameter)
        {
            if (parameter.GetType() == typeof(String) && parameter.Equals("Login"))
            {
                ((App)App.Current).changeViewModel(new LoginViewModel());
            }
            else
            {
                IPasswordProvider passwordProv = parameter as IPasswordProvider;
                CreateAccountModel.Password = PasswordUtils.ConvertToUnsecureString(passwordProv.Password);
                Create_Account();
            }
        }

        public void Update() { }

    }
}