function getRandomString () {
    return Math.random().toString(36).substring(2, 15)
}
  
async function checkIsOnline () {
    if (!window.navigator.onLine) return false;
  
    // avoid CORS errors with a request to your own origin
    let url = new URL(window.location.origin);
    if (window.location.origin.includes("127.0.0.1") || window.location.origin.includes("localhost")){
        url = new URL("https://cors-test.appspot.com/test");
    }
  
    // random value to prevent cached responses
    url.searchParams.set('rand', getRandomString())
  
    try {
      const response = await fetch(
        url.toString(),
        { method: 'HEAD' },
      )
  
      return true
    } catch (e) {
        console.log("Caught rand failure")
        console.log(e);
        return false
    }
}

const ConnectionCheck = () => ({
    isOnline: null,
    async init() {
        this.isOnline = await checkIsOnline(); 
    },
    async handleClick(onlineStatusCallback) {
        this.isOnline = null; 
        this.isOnline = await checkIsOnline();
        onlineStatusCallback(this.isOnline);
        if (this.isOnline) {
            alert("You are online! Please consider disabling your internet connection while using Holocron to reduce chance of secret key leakage");
        } else {
            alert("You are offline");
        }
    },

    bind: {
        ['x-html']() { return /*html*/`
        <button 
            x-init="handleClick(onlineStatusCallback)",
            @click="handleClick(onlineStatusCallback)" 
            x-text="isOnline === null ? 'LOADING...' : isOnline ? 'ONLINE' : 'OFFLINE'"
            :class="isOnline === null ? loadingButtonClass : isOnline === true ? onlineButtonClass : offlineButtonClass",
            type="button"
            >
        </button
        `},
    },
});

export default ConnectionCheck;