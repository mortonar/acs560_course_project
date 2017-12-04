using System;

namespace BookTracker.Models
{
    public class CreateAccountModel
    {
        private String _userName;
        private String _email;
        private String _password;

        public CreateAccountModel()
        {
            _userName = "";
            _email = "";
            _password = "";
        }

        public String UserName
        {
            get { return _userName; }
            set { _userName = value; }
        }

        public String Email
        {
            get { return _email; }
            set { _email = value; }
        }

        public String Password
        {
            get { return _password; }
            set { _password = value; }
        }
    }
}