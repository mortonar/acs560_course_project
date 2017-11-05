using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace BookTracker.Messaging.Requests
{
    public class Base
    {
        public string Token { get; set; }
        public string Action { get; set; }
        public object Payload { get; set; }
    }
}
