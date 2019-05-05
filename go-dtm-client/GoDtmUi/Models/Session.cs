using Newtonsoft.Json;



namespace GoDtmUI.Models
{
    public class Session
    {
        [JsonProperty("token")]
        public string Token { get; set; }

        [JsonProperty("user")]
        public User User { get; set; }

        [JsonProperty("expired")]
        public int Expired { get; set; }


        public static Session FromJson(string json) => JsonConvert.DeserializeObject<Session>(json, Converter.Settings);
    }
}