using Newtonsoft.Json;



namespace GoDtmUI.Models
{
    public class TypesInfo
    {
        [JsonProperty("types")]
        public Status[] Statuses { get; set; }
        
        
        public static TypesInfo FromJson(string json) => JsonConvert.DeserializeObject<TypesInfo>(json, Converter.Settings);
    }
}