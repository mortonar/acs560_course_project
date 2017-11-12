using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace BookTracker.Messaging.Responses
{
    public class Base
    {
        public Boolean Success { get; set; }

        public string Status { get; set; }

        public object Payload { get; set; }
    }
}
