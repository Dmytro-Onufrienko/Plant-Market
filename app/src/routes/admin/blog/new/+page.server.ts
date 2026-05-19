import { fail, redirect } from "@sveltejs/kit";
import { admin } from '$lib/modules/admin/api';
import { blogPostSchema } from "$lib/server/validations/schemas";
import type { Actions } from "./$types";

export const actions: Actions = {
  default: async ({ request }) => {
    const cookie = request.headers.get("cookie") ?? undefined;
    const form = await request.formData();

    let coverImage: string | null = null;
    const file = form.get("coverImage") as File;
    if (file && typeof file === "object" && "size" in file && file.size > 0) {
      try {
        const fd = new FormData();
        fd.append("image", file);
        const result = await admin.plants.upload(fd, cookie);
        coverImage = result.url;
      } catch (e: any) {
        console.error("Image upload error:", e);
        return fail(500, {
          error: "Image upload error: " + (e?.message || String(e)),
        });
      }
    }

    const raw = {
      title: String(form.get("title") ?? ""),
      slug: String(form.get("slug") ?? ""),
      content: String(form.get("content") ?? ""),
      excerpt: String(form.get("excerpt") ?? ""),
      published: form.get("published") === "on",
    };

    const parsed = blogPostSchema.safeParse(raw);
    if (!parsed.success) {
      return fail(400, {
        errors: parsed.error.flatten().fieldErrors,
        values: raw,
      });
    }

    try {
      await admin.blog.create(
        {
          ...parsed.data,
          cover_image: coverImage,
        },
        cookie,
      );
    } catch {
      return fail(500, { error: "Failed to create post" });
    }

    redirect(302, "/admin/blog");
  },
};
