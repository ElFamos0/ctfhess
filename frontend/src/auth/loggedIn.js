import { getRequest } from "@/requests/getRequest";

async function loggedIn() {
    let response = await getRequest("/login/loggedin", "GET");
    let data = await response.data
    return data;
}

export { loggedIn };