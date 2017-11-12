using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace BookTracker.Messaging.Requests
{
    public class BookSearch
    {
        public string Author { get; set; }
        public string Title { get; set; }
    }
}
