export function getCookie(cname: string): string {
    let name = cname + "=";
    let decodedCookie = decodeURIComponent(document.cookie);
    let ca = decodedCookie.split(';');
    for(let i = 0; i <ca.length; i++) {
        let c = ca[i];
        while (c.charAt(0) == ' ') {
            c = c.substring(1);
        }
        if (c.indexOf(name) == 0) {
            return c.substring(name.length, c.length);
        }
    }
    return "";
}

export function setCookie(cname: string, cvalue: string, ex: string) {
    const d = new Date(ex);
    //let expires = "expires="+ d.toUTCString();
    //document.cookie = cname + "=" + cvalue + ";" + expires + ";path=/";
    document.cookie = cname + "=" + cvalue + ";path=/";
}