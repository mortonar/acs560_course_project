using System;

namespace BookTracker.Messaging.Responses
{
    public class Base
    {
        public Boolean Success { get; set; }

        public string Status { get; set; }

        public object Payload { get; set; }
    }
}
