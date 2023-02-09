<script>
    import { enhance } from "$app/forms";

    import * as webauthnJson from "@github/webauthn-json/browser-ponyfill"

    import * as storage from "$lib/js/storage.js"

    import TextInput from "$lib/components/inputs/TextInput.svelte"

    let submitError

    async function webauthnRegisterChallenge({ data, cancel }) {
        submitError = undefined

        const name = data.get("name")
        const authToken = await storage.getAuthToken()

        const response = await fetch(`/api/v1/webauthn/register`, {
            method: "GET",
            headers: {
                "Content-type": "application/json",
                "Authorization": authToken,
            },
        })
        const json = await response.json()

        if (response.status !== 200) {
            submitError = json.error.code
        } else if (response.status === 200) {
            const webauthnChallengeId = json.webauthnChallengeId
            const credential = await webauthnJson.create(webauthnJson.parseCreationOptionsFromJSON(json.options))

            const finishResponse = await fetch("/api/v1/webauthn/register", {
                method: "POST",
                headers: {
                    "Content-type": "application/json",
                    "Authorization": authToken,
                },
                body: JSON.stringify({
                    webauthnChallengeId: webauthnChallengeId,
                    name: name,
                    credential: credential,
                })
            })
            const finishJson = await finishResponse.json()

            if (finishResponse.status !== 200) {
                submitError = finishJson.error.code
            } else if (finishResponse.status === 200) {
                console.log(finishJson)
                console.log("reigsted")
            }
        }

        cancel()
    }
</script>

<form method="POST" use:enhance={webauthnRegisterChallenge}>
    <TextInput classList="flex-grow" label="Name" name="name" description="Enter a namae."/>

    <button type="submit">start</button>
</form>