import { z } from "zod";
import { useForm } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
import {
  Form,
  FormControl,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from "@/components/ui/form";
import { Input } from "@/components/ui/input";
import { Button } from "@/components/ui/button";
import { Textarea } from "@/components/ui/textarea";
import { useState, useRef, useEffect } from "react";
import { CameraIcon } from "lucide-react";
import { Profile } from "@/interface/Profile";
import { toast } from "sonner";
import axios from "axios";
import api from "@/lib/axios";

const formSchema = z.object({
  name: z.string().min(1, "Name is required"),
  phoneNumber: z
    .string()
    .length(10, "Phone number must be 10 digits")
    .optional(),
  address: z.string().min(1, "Address is required").optional(),
  description: z.string().min(1, "Description is required").optional(),
  profileImage: z.any().optional(),
});

type ProfileFormValues = z.infer<typeof formSchema>;

export default function ProfileForm() {
  const [preview, setPreview] = useState<string | null>(null);
  const [profile, setProfile] = useState<Profile>();
  const fileInputRef = useRef<HTMLInputElement>(null);

  const form = useForm<ProfileFormValues>({
    resolver: zodResolver(formSchema),
    defaultValues: {
      name: "",
      phoneNumber: "",
      address: "",
      description: "",
      profileImage: undefined,
    },
  });

  const handleImageChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const file = e.target.files?.[0];
    if (file) {
      setPreview(URL.createObjectURL(file));
      form.setValue("profileImage", file);
    }
  };

  const handleAvatarClick = () => {
    fileInputRef.current?.click();
  };

  const onSubmit = async (values: ProfileFormValues) => {
    try {
      const res = await api.put("/user/profile", values, {
        headers: {
          "Content-Type": "multipart/form-data",
        },
      });
      if (res.status === 200) {
        setProfile(res.data.data);
        toast.success("Profile updated successfully!");
      }
    } catch (error) {
      if (axios.isAxiosError(error)) {
        toast.error(error.response?.data.error || "An unknown error occured!");
      }
    }
  };

  useEffect(() => {
    async function fetchProfile() {
      try {
        const res = await api.get("/user/profile");
        if (res.status === 200) {
          setProfile(res.data.data);
          setPreview(res.data.data.avatarUrl);
        }
      } catch (error) {
        if (axios.isAxiosError(error)) {
          toast.error(error.response?.data.error || "An error occured");
        }
      }
    }

    fetchProfile();
  }, []);

  useEffect(() => {
    if (profile) {
      form.reset({
        name: profile.name,
        phoneNumber: profile.phoneNumber || "",
        address: profile.address || "",
        description: profile.description || "",
        profileImage: undefined,
      });
    }
  }, [profile, form]);

  return (
    <div className="w-5/8 mx-auto p-6 space-y-6 min-h-screen">
      <h2 className="text-2xl font-semibold">Account Information</h2>
      <div className="flex justify-center">
        <div
          className="relative w-24 h-24 rounded-full bg-gray-200 flex items-center justify-center cursor-pointer overflow-hidden group"
          onClick={handleAvatarClick}
        >
          {preview ? (
            <img
              src={preview}
              alt="Profile Preview"
              className="w-full h-full object-cover"
            />
          ) : (
            <span className="text-gray-500 text-3xl">👤</span>
          )}

          <div className="absolute inset-0 bg-black bg-opacity-50 flex items-center justify-center opacity-0 group-hover:opacity-100 transition-opacity">
            <CameraIcon className="text-white w-6 h-6" />
          </div>

          <input
            type="file"
            accept="image/*"
            ref={fileInputRef}
            onChange={handleImageChange}
            className="hidden"
          />
        </div>
      </div>

      <Form {...form}>
        <form onSubmit={form.handleSubmit(onSubmit)} className="space-y-5">
          <FormField
            control={form.control}
            name="name"
            render={({ field }) => (
              <FormItem>
                <FormLabel>
                  Name<span className="text-red-500">*</span>
                </FormLabel>
                <FormControl>
                  <Input placeholder="Enter your name" {...field} />
                </FormControl>
                <FormMessage />
              </FormItem>
            )}
          />

          <FormField
            control={form.control}
            name="phoneNumber"
            render={({ field }) => (
              <FormItem>
                <FormLabel>Phone Number</FormLabel>
                <FormControl>
                  <Input placeholder="Enter your phone number" {...field} />
                </FormControl>
                <FormMessage />
              </FormItem>
            )}
          />

          <FormField
            control={form.control}
            name="address"
            render={({ field }) => (
              <FormItem>
                <FormLabel>Address</FormLabel>
                <FormControl>
                  <Input placeholder="Enter your address" {...field} />
                </FormControl>
                <FormMessage />
              </FormItem>
            )}
          />

          <FormField
            control={form.control}
            name="description"
            render={({ field }) => (
              <FormItem>
                <FormLabel>Description</FormLabel>
                <FormControl>
                  <Textarea
                    placeholder="Enter a short bio or description"
                    {...field}
                  />
                </FormControl>
                <FormMessage />
              </FormItem>
            )}
          />

          <div className="w-full flex justify-center">
            <Button
              type="submit"
              className="w-2/12 bg-yellow-400 hover:bg-yellow-500 text-white"
            >
              Save
            </Button>
          </div>
        </form>
      </Form>
    </div>
  );
}
