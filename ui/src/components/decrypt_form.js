const DecryptForm = () => ({
    cypher: "",
    ascertainment: "",
    decodedMessage: "",
    isCopied: false,
    handleClick() { 
        this.isCopied = false;
        if (this.decodedMessage === '') {
            results = Alpine.store('global_funcs').go.decrypt(
                this.ascertainment,
                this.cypher
                
            );
            if (results.error != '') {
                console.log(results);
                alert(results.error);
            } else {
                this.decodedMessage = results.data.decoded;
            }
        } else {
            this.decodedMessage = "";
            this.ascertainment = "";
        }
    },
    isSubmitable() {
        return (this.decodedMessage != "") ||
        (
            this.cypher != "" &&
            this.ascertainment != ""
        );
    },
    handleCopy() {
        navigator.clipboard.writeText(this.decodedMessage);
        this.isCopied = true;
    },
    bind: {
        ['x-html']() { return /*html*/`
        <section>
        <form >
            <input type="text" x-model="cypher" placeholder="Encoded Holocron Cypher">
            <input  type="password" x-model="ascertainment" placeholder="Correct Answer to Prompt (ex. Open Sesame)">
            <button 
                :disabled="isSubmitable() === false" 
                :class="decodedMessage === '' ? init_button_class : submitted_button_class", 
                @click="handleClick()" 
                x-text="decodedMessage === '' ? 'SIGN' : 'RESET'"
                type="button"
            >
            
            </button>

        </form>
        </section>
        <section>


        <template x-if="decodedMessage != ''">
        <figure>
			<div class="content">
				<h3 class="title">Decrypted Holocron Message</h3>
                <div class="inputs">
                    <input type="text" class="--rounded-left-full" x-model="decodedMessage" class="addon --rounded-right-full" aria-label="copy" readonly>
                    <button @click="handleCopy()">
                        <img height="25" :src="isCopied ? $store.images.svgs.check : $store.images.svgs.copy">
                    </button
                </div>
            </div>
        </figure>
        </template>
        <blockquote class="--family-sans" cite="Dr. Otto Octavius (2004)">
        <p>Being brilliant is not enough...  You have to work hard.  Intelligence is not a privilege, it's a gift, and you use it for the good of mankind.</p>
        </blockquote>
        </section>
        `},
    },
});

export default DecryptForm;