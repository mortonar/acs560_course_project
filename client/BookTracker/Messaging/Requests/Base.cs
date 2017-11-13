namespace BookTracker.Messaging.Requests
{
    public class Base
    {
        public string Token { get; set; }
        public string Action { get; set; }
        public object Payload { get; set; }
    }
}
