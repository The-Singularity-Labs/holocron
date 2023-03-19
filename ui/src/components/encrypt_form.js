const EncryptForm = () => ({
    name: "",
    gatekeeper: "",
    ascertainment: "",
    name: "",
    treasure: "",
    salt: "07121994",
    signedMessage: "",
    signedImage: "",
    isCopied: false,
    handleClick() { 
        this.isCopied = false;
        if (this.signedMessage === '') {
            results = Alpine.store('global_funcs').go.encrypt(
                this.name,
                this.gatekeeper,
                this.ascertainment,
                this.treasure,
                this.salt
            );
            if (results.error != '') {
                console.log(results);
                alert(results.error);
            } else {
                this.signedMessage = results.data.encrypted;
                this.signedImage = results.data.base64_qr_code;
            }
        } else {
            this.signedMessage = "";
            this.signedImage = "";
            this.ascertainment = "";
        }
    },
    isSubmitable() {
        return (this.signedMessage != "") ||
        (
            this.name != "" &&
            this.gatekeeper != "" &&
            this.ascertainment != "" &&
            this.treasure != "" &&
            this.salt != ""
        );
    },
    handleCopy() {
        navigator.clipboard.writeText(this.signedMessage);
        this.isCopied = true;
    },
    bind: {
        ['x-html']() { return /*html*/`
        <section>
        <form >
            <input type="text" x-model="name" placeholder="Holocron Name">
            <input  type="text" x-model="gatekeeper" placeholder="Prompt (ex. Secret of the Universe?)">
            <input  type="password" x-model="ascertainment" placeholder="Correct Anser (ex 4)">
            <textarea  type="text" x-model="treasure" placeholder="Message to encrypt"></textarea>
            <input  type="text" x-model="salt" placeholder="Salt">
            <button 
                :disabled="isSubmitable() === false" 
                :class="signedMessage === '' ? init_button_class : submitted_button_class", 
                @click="handleClick()" 
                x-text="signedMessage === '' ? 'SIGN' : 'RESET'"
                type="button"
            >
            
            </button>

        </form>
        </section>
        <section>


        <template x-if="signedMessage != ''">
        <figure>
            <img :style="{'max-width': '480px', 'width': '100%', 'margin-left': 'auto', 'margin-right': 'auto' }" :src="signedImage" alt="QR Code of Signed Message">
            <figcaption>QR Code of Encrypted Holocron</figcaption>
			<div class="content">
				<h3 class="title">Encrypted Holocron</h3>
                <div class="inputs">
                    <input type="text" class="--rounded-left-full" x-model="signedMessage" class="addon --rounded-right-full" aria-label="copy" readonly>
                    <button @click="handleCopy()">
                        <img height="25" :src="isCopied ? $store.images.svgs.check : $store.images.svgs.copy">
                    </button
                </div>
            </div>
        </figure>
        </template>
        <blockquote class="--family-sans" cite="Robin Morgan (2014)">
        <p>Knowledge is power. Information is power. The secreting or hoarding of knowledge or information may be an act of tyranny camouflaged as humility.</p>
        </blockquote>
        </section>
        `},
    },
});

export default EncryptForm;