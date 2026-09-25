create index if not exists scim_users_directory_id_id_idx
    on scim_users (scim_directory_id, id);

create index if not exists scim_user_group_memberships_group_id_user_id_idx
    on scim_user_group_memberships (scim_group_id, scim_user_id);
