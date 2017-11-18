namespace BookTracker.Messaging.Requests
{
    public class Login
    {
        public string UserName { get; set; }
        public string Email { get; set; }
        public string EncryptedPass { get; set; }
    }
}
