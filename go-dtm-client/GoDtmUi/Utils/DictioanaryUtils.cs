using System.Collections.Generic;



namespace GoDtmUI.Utils
{
    public static class DictioanaryUtils
    {
        public static TValue GetValueAlways<TKey, TValue>(this Dictionary<TKey, TValue> dict, TKey key, TValue defaultValue = default)
        {
            if (key == null || dict == null || !dict.ContainsKey(key))
                return default;

            return dict[key];
        }
    }
}