using System.Security;

namespace BookTracker.HelperClasses
{
    public interface IPasswordProvider
    {
        SecureString Password { get; }
    }
}
