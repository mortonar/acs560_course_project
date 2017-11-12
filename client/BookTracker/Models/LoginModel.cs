using System;

namespace BookTracker.Models
{
    public class LoginModel
    {
        private String _userName;
        private String _password;

        public LoginModel()
        {
            _userName = "BookTracker";
            _password = "booktracker";
        }

        public String UserName
        {
            get { return _userName; }
            set { _userName = value; }
        }

        public String Password
        {
            get { return _password; }
            set { _password = value; }
        }
    }
}