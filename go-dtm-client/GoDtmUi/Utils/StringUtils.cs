using System;

namespace GoDtmUI.Utils
{
    public static class StringUtils
    {
        public static bool IsNullOrEmpty(this string str)
        {
            return String.IsNullOrEmpty(str);
        }
    }
}