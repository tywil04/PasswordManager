import { redirect } from "@sveltejs/kit";

// redirect root to /home
export function load() {
    throw redirect(302, "/home")
}