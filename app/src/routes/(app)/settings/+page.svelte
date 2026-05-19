<script lang="ts">
    import { enhance } from "$app/forms";
    import type { PageData, ActionData } from "./$types";

    let { data, form }: { data: PageData; form: any } = $props();
</script>

<svelte:head>
    <title>Settings — Monstera Shop</title>
</svelte:head>

<div class="bg-green-900 py-10">
    <div class="container">
        <p
            class="text-green-400 text-xs font-semibold tracking-widest uppercase mb-1"
        >
            Your account
        </p>
        <h1 class="text-3xl font-bold text-white">Settings</h1>
    </div>
</div>

<div class="container py-10 max-w-xl space-y-6">
    <!-- Profile info (read-only) -->
    <div class="bg-white rounded-2xl border border-stone-100 shadow-sm p-6">
        <div class="flex items-center gap-4">
            <div
                class="w-14 h-14 rounded-full bg-green-700 flex items-center justify-center text-white text-xl font-bold"
            >
                {data.user.name[0].toUpperCase()}
            </div>
            <div>
                <p class="font-bold text-stone-900">{data.user.name}</p>
                <p class="text-sm text-stone-400">{data.user.email}</p>
                {#if data.user.role === "admin"}
                    <span
                        class="mt-1 inline-block px-2 py-0.5 rounded-full bg-bark-100 text-bark-700 text-xs font-semibold"
                        >Admin</span
                    >
                {/if}
            </div>
        </div>
    </div>

    <!-- Change email -->
    <div class="bg-white rounded-2xl border border-stone-100 shadow-sm p-6">
        <h2 class="text-base font-bold text-stone-900 mb-4">Change Email</h2>

        {#if form?.action === "changeEmail" && form?.success}
            <p class="text-sm text-green-700 font-medium mb-3">
                ✓ Email updated successfully. Please log in again with your new
                email.
            </p>
        {/if}

        <form
            method="POST"
            action="?/changeEmail"
            use:enhance
            class="space-y-4"
        >
            <div>
                <label
                    for="newEmail"
                    class="block text-sm font-medium text-stone-700 mb-1"
                    >New Email *</label
                >
                <input
                    id="newEmail"
                    name="newEmail"
                    type="email"
                    value={form?.action === "changeEmail"
                        ? (form?.values?.newEmail ?? "")
                        : ""}
                    placeholder="new@example.com"
                    class="w-full px-4 py-2.5 rounded-xl border text-sm outline-none
						{form?.action === 'changeEmail' && form?.errors?.newEmail
                        ? 'border-red-400 bg-red-50'
                        : 'border-stone-200 focus:border-green-500'}"
                />
                {#if form?.action === "changeEmail" && form?.errors?.newEmail}
                    <p class="mt-1 text-xs text-red-500">
                        {form.errors.newEmail[0]}
                    </p>
                {/if}
            </div>
            <div>
                <label
                    for="emailPass"
                    class="block text-sm font-medium text-stone-700 mb-1"
                    >Current Password *</label
                >
                <input
                    id="emailPass"
                    name="currentPassword"
                    type="password"
                    placeholder="Your current password"
                    class="w-full px-4 py-2.5 rounded-xl border text-sm outline-none
						{form?.action === 'changeEmail' && form?.errors?.currentPassword
                        ? 'border-red-400 bg-red-50'
                        : 'border-stone-200 focus:border-green-500'}"
                />
                {#if form?.action === "changeEmail" && form?.errors?.currentPassword}
                    <p class="mt-1 text-xs text-red-500">
                        {form.errors.currentPassword[0]}
                    </p>
                {/if}
            </div>
            <button
                type="submit"
                class="px-5 py-2.5 rounded-xl bg-green-700 text-white text-sm font-semibold hover:bg-green-600 active:scale-95 transition-all"
            >
                Update Email
            </button>
        </form>
    </div>

    <!-- Change password -->
    <div class="bg-white rounded-2xl border border-stone-100 shadow-sm p-6">
        <h2 class="text-base font-bold text-stone-900 mb-4">Change Password</h2>

        {#if form?.action === "changePassword" && form?.success}
            <p class="text-sm text-green-700 font-medium mb-3">
                ✓ Password updated successfully.
            </p>
        {/if}

        <form
            method="POST"
            action="?/changePassword"
            use:enhance
            class="space-y-4"
        >
            <div>
                <label
                    for="curPass"
                    class="block text-sm font-medium text-stone-700 mb-1"
                    >Current Password *</label
                >
                <input
                    id="curPass"
                    name="currentPassword"
                    type="password"
                    placeholder="Your current password"
                    class="w-full px-4 py-2.5 rounded-xl border text-sm outline-none
						{form?.action === 'changePassword' && form?.errors?.currentPassword
                        ? 'border-red-400 bg-red-50'
                        : 'border-stone-200 focus:border-green-500'}"
                />
                {#if form?.action === "changePassword" && form?.errors?.currentPassword}
                    <p class="mt-1 text-xs text-red-500">
                        {form.errors.currentPassword[0]}
                    </p>
                {/if}
            </div>
            <div>
                <label
                    for="newPass"
                    class="block text-sm font-medium text-stone-700 mb-1"
                    >New Password *</label
                >
                <input
                    id="newPass"
                    name="newPassword"
                    type="password"
                    placeholder="Min 8 characters"
                    class="w-full px-4 py-2.5 rounded-xl border text-sm outline-none
						{form?.action === 'changePassword' && form?.errors?.newPassword
                        ? 'border-red-400 bg-red-50'
                        : 'border-stone-200 focus:border-green-500'}"
                />
                {#if form?.action === "changePassword" && form?.errors?.newPassword}
                    <p class="mt-1 text-xs text-red-500">
                        {form.errors.newPassword[0]}
                    </p>
                {/if}
            </div>
            <div>
                <label
                    for="confPass"
                    class="block text-sm font-medium text-stone-700 mb-1"
                    >Confirm New Password *</label
                >
                <input
                    id="confPass"
                    name="confirmPassword"
                    type="password"
                    placeholder="Repeat new password"
                    class="w-full px-4 py-2.5 rounded-xl border text-sm outline-none
						{form?.action === 'changePassword' && form?.errors?.confirmPassword
                        ? 'border-red-400 bg-red-50'
                        : 'border-stone-200 focus:border-green-500'}"
                />
                {#if form?.action === "changePassword" && form?.errors?.confirmPassword}
                    <p class="mt-1 text-xs text-red-500">
                        {form.errors.confirmPassword[0]}
                    </p>
                {/if}
            </div>
            <button
                type="submit"
                class="px-5 py-2.5 rounded-xl bg-green-700 text-white text-sm font-semibold hover:bg-green-600 active:scale-95 transition-all"
            >
                Update Password
            </button>
        </form>
    </div>

    <!-- Mock payment card (UI only, user only) -->
    {#if data.user.role !== "admin"}
        <div class="bg-white rounded-2xl border border-stone-100 shadow-sm p-6">
            <h2 class="text-base font-bold text-stone-900 mb-1">
                Saved Payment Method
            </h2>
            <p class="text-xs text-stone-400 mb-4">
                Demo mode — card details are not stored
            </p>

            <!-- Card preview -->
            <div
                class="relative w-full max-w-xs h-40 rounded-2xl bg-gradient-to-br from-green-700 to-green-900 p-5 text-white mb-4 overflow-hidden"
            >
                <div class="absolute top-4 right-4 opacity-30">
                    <svg width="48" height="48" viewBox="0 0 48 48" fill="none">
                        <circle cx="18" cy="24" r="14" fill="white" />
                        <circle
                            cx="30"
                            cy="24"
                            r="14"
                            fill="white"
                            opacity="0.7"
                        />
                    </svg>
                </div>
                <p class="text-xs opacity-60 mb-6">MONSTERA CARD</p>
                <p class="font-mono text-base tracking-widest mb-3">
                    •••• •••• •••• 4242
                </p>
                <div class="flex justify-between text-xs opacity-70">
                    <span>VALID THRU 12/28</span>
                    <span>DEMO ONLY</span>
                </div>
            </div>

            <div class="grid grid-cols-2 gap-4">
                <div>
                    <label
                        for="cardNum"
                        class="block text-sm font-medium text-stone-700 mb-1"
                        >Card Number</label
                    >
                    <input
                        id="cardNum"
                        type="text"
                        maxlength="19"
                        placeholder="4242 4242 4242 4242"
                        class="w-full px-4 py-2.5 rounded-xl border border-stone-200 text-sm outline-none focus:border-green-500 font-mono"
                    />
                </div>
                <div>
                    <label
                        for="cardName"
                        class="block text-sm font-medium text-stone-700 mb-1"
                        >Cardholder Name</label
                    >
                    <input
                        id="cardName"
                        type="text"
                        placeholder="Sofia Kovalenko"
                        class="w-full px-4 py-2.5 rounded-xl border border-stone-200 text-sm outline-none focus:border-green-500"
                    />
                </div>
                <div>
                    <label
                        for="cardExp"
                        class="block text-sm font-medium text-stone-700 mb-1"
                        >Expiry</label
                    >
                    <input
                        id="cardExp"
                        type="text"
                        maxlength="5"
                        placeholder="MM/YY"
                        class="w-full px-4 py-2.5 rounded-xl border border-stone-200 text-sm outline-none focus:border-green-500 font-mono"
                    />
                </div>
                <div>
                    <label
                        for="cardCvc"
                        class="block text-sm font-medium text-stone-700 mb-1"
                        >CVC</label
                    >
                    <input
                        id="cardCvc"
                        type="text"
                        maxlength="3"
                        placeholder="123"
                        class="w-full px-4 py-2.5 rounded-xl border border-stone-200 text-sm outline-none focus:border-green-500 font-mono"
                    />
                </div>
            </div>

            <div
                class="mt-4 p-3 rounded-xl bg-amber-50 border border-amber-100 text-xs text-amber-700"
            >
                ⚠️ This is a demo feature. No real card data is saved or
                processed.
            </div>
        </div>
    {/if}
</div>
