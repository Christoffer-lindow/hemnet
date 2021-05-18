import axios from "axios"

export const getProperties = async (city) => {
    try {
        const result = await axios.get(`http://localhost:8082/api/properties/${city}`)
        return result
    } catch (ex) {
        return {"error" : ex}
    }
}