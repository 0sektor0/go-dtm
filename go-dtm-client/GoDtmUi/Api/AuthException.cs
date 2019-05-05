using System;

namespace GoDtmUI.Api
{
    public class AuthException : Exception
    {
        public AuthException() : base("wrong login or password")
        {
            
        }
    }
}