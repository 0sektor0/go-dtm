using System;
using GoDtmUI.Models;



namespace GoDtmUI.Api
{
    public static class Utils
    {
        public static bool IsExpired(this Session session)
        {
            var isExpired = session.Expired < GetUnixTimestamp();
            return isExpired;
        }

        public static int GetUnixTimestamp()
        {
            var unixTime = (int)(DateTime.UtcNow - new DateTime(1970, 1, 1)).TotalSeconds;
            return unixTime;
        }
    }
}