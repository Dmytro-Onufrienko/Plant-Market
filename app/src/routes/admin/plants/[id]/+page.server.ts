import { fail, redirect, error } from "@sveltejs/kit";
import { ApiError } from '$lib/modules/common/api/fetcher';
import { categories as categoriesApi } from '$lib/modules/plants/api';
import { admin } from '$lib/modules/admin/api';
import { plantSchema } from "$lib/server/validations/schemas";
import type { PageServerLoad, Actions } from "./$types";

export const load: PageServerLoad = async ({ params, request }) => {
  const cookie = request.headers.get("cookie") ?? undefined;
  const plantId = Number(params.id);

  let plantData;
  try {
    // Use by-id to get plant for editing
    const res = await admin.plants.list(cookie);
    const found = res.plants.find((p) => p.id === plantId);
    if (!found) error(404, "Plant not found");
    plantData = found;
  } catch (e) {
    if (e instanceof ApiError && e.status === 404)
      error(404, "Plant not found");
    throw e;
  }

  const catsData = await categoriesApi.list(cookie);

  return {
    plant: plantData,
    categories: catsData.categories,
  };
};

export const actions: Actions = {
  default: async ({ request, params }) => {
    const cookie = request.headers.get("cookie") ?? undefined;
    const plantId = Number(params.id);
    const form = await request.formData();

    // Images to keep (sent as hidden inputs)
    const keepImages = form.getAll("keepImage") as string[];

    // New uploads → Go upload endpoint
    const newFiles = form.getAll("images") as File[];
    const newPaths: string[] = [];
    for (const file of newFiles) {
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
          newPaths.push(result.url);
        } catch (e: any) {
          console.error("Image upload error:", e);
          return fail(500, {
            error: "Image upload error: " + (e?.message || String(e)),
          });
        }
      }
    }

    const finalImages = [...keepImages, ...newPaths];

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
      await admin.plants.update(
        plantId,
        {
          name: parsed.data.name,
          slug: parsed.data.slug,
          description: parsed.data.description ?? "",
          price: parsed.data.price,
          category_id: parsed.data.categoryId ?? null,
          stock: parsed.data.stock,
          featured: parsed.data.featured,
          images: finalImages,
        },
        cookie,
      );
    } catch {
      return fail(500, { error: "Failed to update plant" });
    }

    redirect(302, "/admin/plants");
  },
};
