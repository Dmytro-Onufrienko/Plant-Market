import { fail, redirect, error } from "@sveltejs/kit";
import { ApiError } from '$lib/modules/common/api/fetcher';
import { admin } from '$lib/modules/admin/api';
import { blogPostSchema } from "$lib/server/validations/schemas";
import type { PageServerLoad, Actions } from "./$types";

export const load: PageServerLoad = async ({ params, request }) => {
  const cookie = request.headers.get("cookie") ?? undefined;
  const data = await admin.blog.list(cookie);
  const post = data.posts.find((p) => p.id === Number(params.id));
  if (!post) error(404, "Post not found");

  return {
    post: {
      id: post.id,
      title: post.title,
      slug: post.slug,
      content: post.content,
      excerpt: post.excerpt,
      coverImage: post.cover_image,
      published: post.published,
      createdAt: post.created_at,
    },
  };
};

export const actions: Actions = {
  default: async ({ request, params }) => {
    const cookie = request.headers.get("cookie") ?? undefined;
    const form = await request.formData();

    let coverImage: string | null = form.get("existingCoverImage") as
      | string
      | null;

    if (
      form.has("removeCoverImage") &&
      form.get("removeCoverImage") === "true"
    ) {
      coverImage = null;
    }

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
      await admin.blog.update(
        Number(params.id),
        {
          ...parsed.data,
          cover_image: coverImage,
        },
        cookie,
      );
    } catch {
      return fail(500, { error: "Failed to update post" });
    }

    redirect(302, "/admin/blog");
  },
};
