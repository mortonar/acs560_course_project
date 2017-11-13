using System;
using System.Collections.Generic;
using System.Linq;
using System.Security;
using System.Text;
using System.Threading.Tasks;

namespace BookTracker.HelperClasses
{
    public interface IPasswordProvider
    {
        SecureString Password { get; }
    }
}
