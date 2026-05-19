import { fail, redirect } from "@sveltejs/kit";
import { categories as categoriesApi } from '$lib/modules/plants/api';
import { admin } from '$lib/modules/admin/api';
import { plantSchema } from "$lib/server/validations/schemas";
import type { PageServerLoad, Actions } from "./$types";

export const load: PageServerLoad = async ({ request }) => {
  const cookie = request.headers.get("cookie") ?? undefined;
  const data = await categoriesApi.list(cookie);
  return { categories: data.categories };
};

export const actions: Actions = {
  default: async ({ request }) => {
    const cookie = request.headers.get("cookie") ?? undefined;
    const form = await request.formData();

    // Upload images to Go backend
    const imageFiles = form.getAll("images") as File[];
    const imagePaths: string[] = [];

    for (const file of imageFiles) {
      if (
        file &&
        typeof file === "object" &&
        "size" in file &&
        (file as File).size > 0
      ) {
        try {
          const fd = new FormData();
          fd.append("image", file as File);
          const result = await admin.plants.upload(fd, cookie);
          imagePaths.push(result.url);
        } catch (e: any) {
          console.error("Image upload error:", e);
          return fail(500, {
            error: "Image upload error: " + (e?.message || String(e)),
          });
        }
      }
    }

    const raw = {
      name: String(form.get("name") ?? ""),
      slug: String(form.get("slug") ?? ""),
      description: String(form.get("description") ?? ""),
      price: Math.round(Number(form.get("price") ?? 0) * 100),
      categoryId: form.get("categoryId")
        ? Number(form.get("categoryId"))
        : null,
      stock: Number(form.get("stock") ?? 0),
      featured: form.get("featured") === "on",
    };

    const parsed = plantSchema.safeParse(raw);
    if (!parsed.success) {
      return fail(400, {
        errors: parsed.error.flatten().fieldErrors,
        values: raw,
      });
    }

    try {
      await admin.plants.create(
        {
          name: parsed.data.name,
          slug: parsed.data.slug,
          description: parsed.data.description ?? "",
          price: parsed.data.price,
          category_id: parsed.data.categoryId ?? null,
          stock: parsed.data.stock,
          featured: parsed.data.featured,
          images: imagePaths,
        },
        cookie,
      );
    } catch {
      return fail(500, { error: "Failed to create plant" });
    }

    redirect(302, "/admin/plants");
  },
};
