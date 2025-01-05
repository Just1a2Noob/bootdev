import os
import shutil

from generate_page import generate_pages_recursive

# Change the paths here
dir_path_static = os.path.expanduser(
    "~/Documents/github/bootdev/static_website_generator/static/"
)
dir_path_public = os.path.expanduser(
    "~/Documents/github/bootdev/static_website_generator/public/"
)
dir_path_content = os.path.expanduser(
    "~/Documents/github/bootdev/static_website_generator/content"
)
template_path = os.path.expanduser(
    "~/Documents/github/bootdev/static_website_generator/template.html"
)


def copy_files_recursive(source_dir_path, destination_dir_path):
    if not os.path.exists(destination_dir_path):
        os.mkdir(destination_dir_path)

    for filename in os.listdir(source_dir_path):
        from_path = os.path.join(source_dir_path, filename)
        dest_path = os.path.join(destination_dir_path, filename)
        print(f" * {from_path} -> {dest_path}")

        # Checks if the from_path is a file
        if os.path.isfile(from_path):
            # Then copies it to dest_path
            shutil.copy(from_path, dest_path)
        else:
            # If its not a file recurive call with new from and dest
            copy_files_recursive(from_path, dest_path)


def main():
    print("Deleting public directory...")
    if os.path.exists(dir_path_public):
        shutil.rmtree(dir_path_public)

    print("Copying static files to public directory")
    copy_files_recursive(dir_path_static, dir_path_public)

    print("Generating content...")
    generate_pages_recursive(dir_path_content, template_path, dir_path_public)


main()
