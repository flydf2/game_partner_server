-- 更新图片链接SQL语句
-- 使用global_image_links.md中的图片链接随机更新数据库表中的图片字段

-- 使用game_partner数据库
USE game_partner;

-- 更新用户表的avatar字段（使用较短的链接）
UPDATE game_partner_users SET avatar = SUBSTRING(
    ELT(
        FLOOR(RAND() * 6) + 1,
        'https://lh3.googleusercontent.com/aida-public/AB6AXuBoBTc_5bZuXQ8l_u2UKFazkITvVt5UY-tB83GE9qFMKnbb7Gz7DBuHH11MCcfExFpNociu2AurEP9Lt2NRc9nvSntdZ9hgcWNL_d-0yyyC7bLbO0F8qFUi1FZ_0xgHBG5ZWEfyBs3f5BMl_rBN4SHJoDd3xp76P8kx7eQBwXzcI46GuMySscFwGrnXs_YK9_ArHQEUVcsZUe0o_yRl84Nf4j3WwXor_Xd2gFFDgNuPdbSQuyQiPQkAovGTm7Cek_vM2ZGapACwBM4',
        'https://lh3.googleusercontent.com/aida-public/AB6AXuChoKqMjV8Zopn6t52b1X_D1IHpomYuX0uMCtUOdsEjdUbzmVSeBr-5ZKxEhOUlx1mPSxANNKH4sghnW_76nSPFe5RQDcs0D0146L6oIlgGmuRs6DcX59vrx67krFal1DduP8vpSj5o2giMgdNHYP2pIzqcKEuJfnSxyvoN_fgqJahQCNwzPm37cGoeZUvBWw_zifeh-pvUxLyRR7_8puuGkKr-5w05-1orNhrAYXS-S8PDY6s1vdzatJ05jaJSvNTo7dS88EwkP2M',
        'https://lh3.googleusercontent.com/aida-public/AB6AXuAPWd42wRhAneFPqDBu00-X6zV_3VmT3LwRUFe1rO4eEyOa-IDEPhh1HH7x2WWdyarjgr_4WzatpsdpqaJXJlMRwo9dtnaBbaYDQBj0-oao709scndaRZifLRuzJRgNVcSzd8XsuGJ-4uBbEUbdAlWFQz_6HkFw-AT6QvtHY4sIV4zpurtJ13FIBZJw5DYnKUc0lTPGLItY507UIdcj7nBxJ2WZ7zMMYQN6FvTlEsX9vRQQ51tpjD9-9ylRac-P1jqyPFi7WtKf__I',
        'https://lh3.googleusercontent.com/aida-public/AB6AXuDAkokFDlJJU2ut7ADXOxXgokfLf-nFbve3PmdniZPK8ZjoxO1reJWexxqbxEwDy4thsHQWxI68ySeYsam7yLDI2EWIKoa7no1yZya_-OvKHcnsOU5ukziXluccC3dUHzW8BMFTNlh6bLkgI8bMSDhUTMaTR2jWDsFNRzAKs9RzZ11om4Ki5c9Xj92fIm685lMzUrq9Bk1QuG6y3TXwjzpc6HZVTEFv-aZfUUXt5s020g03CNeOFTD-_F8HtoWP1VEGiCAo5k4gC9o',
        'https://lh3.googleusercontent.com/aida-public/AB6AXuBDI76aOeylc3VRSRNZ04zYnjdju83SPGLpYh8d2uB5bogOv5w6wJcvfx6AJMDR5RQKBUrAlNQptRHNf3XJ61WGVmFOAc3-gFksvrlt6w6F_TaBfQ6Mij2zaM_KjWfcqlVkL82KYbgijsqdy-ptNMAsnQRBVHpqk2dGVjd6UsWwcuOnY9ToQPlAZwJXaT0sdGvDFkoLuAn_IIt1cP9aJJ9LXF0dXPKQK0bAJi57JnH1pgIKL4F5lFuQICzWqRy0VNT1R6BfwuOFvKg',
        'https://lh3.googleusercontent.com/aida-public/AB6AXuBHACwqs1DoKdU4I8vIetAZ1Ncf0TX2gOpmC_WxRtq2BmFEm-uJnLxeY3lrs_zfMradBduqlzQqbwRp5skRm1dQ4Ff85ob-zSZNkEhwrkNdJFgidiWuFw1ZdgMsuWEo0w82df7KisJHze0g2woOiTDZeoS_BJMCUr_VeiFECrJnlibmkGXTF67XclMyzRwRH0OjcUgdGN1GctXt_aRiiUbW4CAPTWV8RwnBLbb_HgnjSqYmbKHcclCaQQmdEbphHpEeCMk8w9JtxY4'
    ), 1, 100
);

-- 更新陪玩专家表的avatar字段（使用较短的链接）
UPDATE game_partner_playmates SET avatar = SUBSTRING(
    ELT(
        FLOOR(RAND() * 6) + 1,
        'https://lh3.googleusercontent.com/aida-public/AB6AXuBoBTc_5bZuXQ8l_u2UKFazkITvVt5UY-tB83GE9qFMKnbb7Gz7DBuHH11MCcfExFpNociu2AurEP9Lt2NRc9nvSntdZ9hgcWNL_d-0yyyC7bLbO0F8qFUi1FZ_0xgHBG5ZWEfyBs3f5BMl_rBN4SHJoDd3xp76P8kx7eQBwXzcI46GuMySscFwGrnXs_YK9_ArHQEUVcsZUe0o_yRl84Nf4j3WwXor_Xd2gFFDgNuPdbSQuyQiPQkAovGTm7Cek_vM2ZGapACwBM4',
        'https://lh3.googleusercontent.com/aida-public/AB6AXuChoKqMjV8Zopn6t52b1X_D1IHpomYuX0uMCtUOdsEjdUbzmVSeBr-5ZKxEhOUlx1mPSxANNKH4sghnW_76nSPFe5RQDcs0D0146L6oIlgGmuRs6DcX59vrx67krFal1DduP8vpSj5o2giMgdNHYP2pIzqcKEuJfnSxyvoN_fgqJahQCNwzPm37cGoeZUvBWw_zifeh-pvUxLyRR7_8puuGkKr-5w05-1orNhrAYXS-S8PDY6s1vdzatJ05jaJSvNTo7dS88EwkP2M',
        'https://lh3.googleusercontent.com/aida-public/AB6AXuAPWd42wRhAneFPqDBu00-X6zV_3VmT3LwRUFe1rO4eEyOa-IDEPhh1HH7x2WWdyarjgr_4WzatpsdpqaJXJlMRwo9dtnaBbaYDQBj0-oao709scndaRZifLRuzJRgNVcSzd8XsuGJ-4uBbEUbdAlWFQz_6HkFw-AT6QvtHY4sIV4zpurtJ13FIBZJw5DYnKUc0lTPGLItY507UIdcj7nBxJ2WZ7zMMYQN6FvTlEsX9vRQQ51tpjD9-9ylRac-P1jqyPFi7WtKf__I',
        'https://lh3.googleusercontent.com/aida-public/AB6AXuDAkokFDlJJU2ut7ADXOxXgokfLf-nFbve3PmdniZPK8ZjoxO1reJWexxqbxEwDy4thsHQWxI68ySeYsam7yLDI2EWIKoa7no1yZya_-OvKHcnsOU5ukziXluccC3dUHzW8BMFTNlh6bLkgI8bMSDhUTMaTR2jWDsFNRzAKs9RzZ11om4Ki5c9Xj92fIm685lMzUrq9Bk1QuG6y3TXwjzpc6HZVTEFv-aZfUUXt5s020g03CNeOFTD-_F8HtoWP1VEGiCAo5k4gC9o',
        'https://lh3.googleusercontent.com/aida-public/AB6AXuBDI76aOeylc3VRSRNZ04zYnjdju83SPGLpYh8d2uB5bogOv5w6wJcvfx6AJMDR5RQKBUrAlNQptRHNf3XJ61WGVmFOAc3-gFksvrlt6w6F_TaBfQ6Mij2zaM_KjWfcqlVkL82KYbgijsqdy-ptNMAsnQRBVHpqk2dGVjd6UsWwcuOnY9ToQPlAZwJXaT0sdGvDFkoLuAn_IIt1cP9aJJ9LXF0dXPKQK0bAJi57JnH1pgIKL4F5lFuQICzWqRy0VNT1R6BfwuOFvKg',
        'https://lh3.googleusercontent.com/aida-public/AB6AXuBHACwqs1DoKdU4I8vIetAZ1Ncf0TX2gOpmC_WxRtq2BmFEm-uJnLxeY3lrs_zfMradBduqlzQqbwRp5skRm1dQ4Ff85ob-zSZNkEhwrkNdJFgidiWuFw1ZdgMsuWEo0w82df7KisJHze0g2woOiTDZeoS_BJMCUr_VeiFECrJnlibmkGXTF67XclMyzRwRH0OjcUgdGN1GctXt_aRiiUbW4CAPTWV8RwnBLbb_HgnjSqYmbKHcclCaQQmdEbphHpEeCMk8w9JtxY4'
    ), 1, 100
);

-- 更新社区帖子表的images字段（随机选择1-3个图片链接，用逗号分隔，限制总长度）
UPDATE game_partner_community_posts SET images = SUBSTRING(
    CONCAT(
        SUBSTRING(
            ELT(
                FLOOR(RAND() * 6) + 1,
                'https://lh3.googleusercontent.com/aida-public/AB6AXuBoBTc_5bZuXQ8l_u2UKFazkITvVt5UY-tB83GE9qFMKnbb7Gz7DBuHH11MCcfExFpNociu2AurEP9Lt2NRc9nvSntdZ9hgcWNL_d-0yyyC7bLbO0F8qFUi1FZ_0xgHBG5ZWEfyBs3f5BMl_rBN4SHJoDd3xp76P8kx7eQBwXzcI46GuMySscFwGrnXs_YK9_ArHQEUVcsZUe0o_yRl84Nf4j3WwXor_Xd2gFFDgNuPdbSQuyQiPQkAovGTm7Cek_vM2ZGapACwBM4',
                'https://lh3.googleusercontent.com/aida-public/AB6AXuChoKqMjV8Zopn6t52b1X_D1IHpomYuX0uMCtUOdsEjdUbzmVSeBr-5ZKxEhOUlx1mPSxANNKH4sghnW_76nSPFe5RQDcs0D0146L6oIlgGmuRs6DcX59vrx67krFal1DduP8vpSj5o2giMgdNHYP2pIzqcKEuJfnSxyvoN_fgqJahQCNwzPm37cGoeZUvBWw_zifeh-pvUxLyRR7_8puuGkKr-5w05-1orNhrAYXS-S8PDY6s1vdzatJ05jaJSvNTo7dS88EwkP2M',
                'https://lh3.googleusercontent.com/aida-public/AB6AXuAPWd42wRhAneFPqDBu00-X6zV_3VmT3LwRUFe1rO4eEyOa-IDEPhh1HH7x2WWdyarjgr_4WzatpsdpqaJXJlMRwo9dtnaBbaYDQBj0-oao709scndaRZifLRuzJRgNVcSzd8XsuGJ-4uBbEUbdAlWFQz_6HkFw-AT6QvtHY4sIV4zpurtJ13FIBZJw5DYnKUc0lTPGLItY507UIdcj7nBxJ2WZ7zMMYQN6FvTlEsX9vRQQ51tpjD9-9ylRac-P1jqyPFi7WtKf__I',
                'https://lh3.googleusercontent.com/aida-public/AB6AXuDAkokFDlJJU2ut7ADXOxXgokfLf-nFbve3PmdniZPK8ZjoxO1reJWexxqbxEwDy4thsHQWxI68ySeYsam7yLDI2EWIKoa7no1yZya_-OvKHcnsOU5ukziXluccC3dUHzW8BMFTNlh6bLkgI8bMSDhUTMaTR2jWDsFNRzAKs9RzZ11om4Ki5c9Xj92fIm685lMzUrq9Bk1QuG6y3TXwjzpc6HZVTEFv-aZfUUXt5s020g03CNeOFTD-_F8HtoWP1VEGiCAo5k4gC9o',
                'https://lh3.googleusercontent.com/aida-public/AB6AXuBDI76aOeylc3VRSRNZ04zYnjdju83SPGLpYh8d2uB5bogOv5w6wJcvfx6AJMDR5RQKBUrAlNQptRHNf3XJ61WGVmFOAc3-gFksvrlt6w6F_TaBfQ6Mij2zaM_KjWfcqlVkL82KYbgijsqdy-ptNMAsnQRBVHpqk2dGVjd6UsWwcuOnY9ToQPlAZwJXaT0sdGvDFkoLuAn_IIt1cP9aJJ9LXF0dXPKQK0bAJi57JnH1pgIKL4F5lFuQICzWqRy0VNT1R6BfwuOFvKg',
                'https://lh3.googleusercontent.com/aida-public/AB6AXuBHACwqs1DoKdU4I8vIetAZ1Ncf0TX2gOpmC_WxRtq2BmFEm-uJnLxeY3lrs_zfMradBduqlzQqbwRp5skRm1dQ4Ff85ob-zSZNkEhwrkNdJFgidiWuFw1ZdgMsuWEo0w82df7KisJHze0g2woOiTDZeoS_BJMCUr_VeiFECrJnlibmkGXTF67XclMyzRwRH0OjcUgdGN1GctXt_aRiiUbW4CAPTWV8RwnBLbb_HgnjSqYmbKHcclCaQQmdEbphHpEeCMk8w9JtxY4'
            ), 1, 100
        ),
        CASE WHEN RAND() > 0.5 THEN CONCAT(',', SUBSTRING(
            ELT(
                FLOOR(RAND() * 6) + 1,
                'https://lh3.googleusercontent.com/aida-public/AB6AXuBoBTc_5bZuXQ8l_u2UKFazkITvVt5UY-tB83GE9qFMKnbb7Gz7DBuHH11MCcfExFpNociu2AurEP9Lt2NRc9nvSntdZ9hgcWNL_d-0yyyC7bLbO0F8qFUi1FZ_0xgHBG5ZWEfyBs3f5BMl_rBN4SHJoDd3xp76P8kx7eQBwXzcI46GuMySscFwGrnXs_YK9_ArHQEUVcsZUe0o_yRl84Nf4j3WwXor_Xd2gFFDgNuPdbSQuyQiPQkAovGTm7Cek_vM2ZGapACwBM4',
                'https://lh3.googleusercontent.com/aida-public/AB6AXuChoKqMjV8Zopn6t52b1X_D1IHpomYuX0uMCtUOdsEjdUbzmVSeBr-5ZKxEhOUlx1mPSxANNKH4sghnW_76nSPFe5RQDcs0D0146L6oIlgGmuRs6DcX59vrx67krFal1DduP8vpSj5o2giMgdNHYP2pIzqcKEuJfnSxyvoN_fgqJahQCNwzPm37cGoeZUvBWw_zifeh-pvUxLyRR7_8puuGkKr-5w05-1orNhrAYXS-S8PDY6s1vdzatJ05jaJSvNTo7dS88EwkP2M',
                'https://lh3.googleusercontent.com/aida-public/AB6AXuAPWd42wRhAneFPqDBu00-X6zV_3VmT3LwRUFe1rO4eEyOa-IDEPhh1HH7x2WWdyarjgr_4WzatpsdpqaJXJlMRwo9dtnaBbaYDQBj0-oao709scndaRZifLRuzJRgNVcSzd8XsuGJ-4uBbEUbdAlWFQz_6HkFw-AT6QvtHY4sIV4zpurtJ13FIBZJw5DYnKUc0lTPGLItY507UIdcj7nBxJ2WZ7zMMYQN6FvTlEsX9vRQQ51tpjD9-9ylRac-P1jqyPFi7WtKf__I',
                'https://lh3.googleusercontent.com/aida-public/AB6AXuDAkokFDlJJU2ut7ADXOxXgokfLf-nFbve3PmdniZPK8ZjoxO1reJWexxqbxEwDy4thsHQWxI68ySeYsam7yLDI2EWIKoa7no1yZya_-OvKHcnsOU5ukziXluccC3dUHzW8BMFTNlh6bLkgI8bMSDhUTMaTR2jWDsFNRzAKs9RzZ11om4Ki5c9Xj92fIm685lMzUrq9Bk1QuG6y3TXwjzpc6HZVTEFv-aZfUUXt5s020g03CNeOFTD-_F8HtoWP1VEGiCAo5k4gC9o',
                'https://lh3.googleusercontent.com/aida-public/AB6AXuBDI76aOeylc3VRSRNZ04zYnjdju83SPGLpYh8d2uB5bogOv5w6wJcvfx6AJMDR5RQKBUrAlNQptRHNf3XJ61WGVmFOAc3-gFksvrlt6w6F_TaBfQ6Mij2zaM_KjWfcqlVkL82KYbgijsqdy-ptNMAsnQRBVHpqk2dGVjd6UsWwcuOnY9ToQPlAZwJXaT0sdGvDFkoLuAn_IIt1cP9aJJ9LXF0dXPKQK0bAJi57JnH1pgIKL4F5lFuQICzWqRy0VNT1R6BfwuOFvKg',
                'https://lh3.googleusercontent.com/aida-public/AB6AXuBHACwqs1DoKdU4I8vIetAZ1Ncf0TX2gOpmC_WxRtq2BmFEm-uJnLxeY3lrs_zfMradBduqlzQqbwRp5skRm1dQ4Ff85ob-zSZNkEhwrkNdJFgidiWuFw1ZdgMsuWEo0w82df7KisJHze0g2woOiTDZeoS_BJMCUr_VeiFECrJnlibmkGXTF67XclMyzRwRH0OjcUgdGN1GctXt_aRiiUbW4CAPTWV8RwnBLbb_HgnjSqYmbKHcclCaQQmdEbphHpEeCMk8w9JtxY4'
            ), 1, 100
        )) ELSE '' END,
        CASE WHEN RAND() > 0.7 THEN CONCAT(',', SUBSTRING(
            ELT(
                FLOOR(RAND() * 6) + 1,
                'https://lh3.googleusercontent.com/aida-public/AB6AXuBoBTc_5bZuXQ8l_u2UKFazkITvVt5UY-tB83GE9qFMKnbb7Gz7DBuHH11MCcfExFpNociu2AurEP9Lt2NRc9nvSntdZ9hgcWNL_d-0yyyC7bLbO0F8qFUi1FZ_0xgHBG5ZWEfyBs3f5BMl_rBN4SHJoDd3xp76P8kx7eQBwXzcI46GuMySscFwGrnXs_YK9_ArHQEUVcsZUe0o_yRl84Nf4j3WwXor_Xd2gFFDgNuPdbSQuyQiPQkAovGTm7Cek_vM2ZGapACwBM4',
                'https://lh3.googleusercontent.com/aida-public/AB6AXuChoKqMjV8Zopn6t52b1X_D1IHpomYuX0uMCtUOdsEjdUbzmVSeBr-5ZKxEhOUlx1mPSxANNKH4sghnW_76nSPFe5RQDcs0D0146L6oIlgGmuRs6DcX59vrx67krFal1DduP8vpSj5o2giMgdNHYP2pIzqcKEuJfnSxyvoN_fgqJahQCNwzPm37cGoeZUvBWw_zifeh-pvUxLyRR7_8puuGkKr-5w05-1orNhrAYXS-S8PDY6s1vdzatJ05jaJSvNTo7dS88EwkP2M',
                'https://lh3.googleusercontent.com/aida-public/AB6AXuAPWd42wRhAneFPqDBu00-X6zV_3VmT3LwRUFe1rO4eEyOa-IDEPhh1HH7x2WWdyarjgr_4WzatpsdpqaJXJlMRwo9dtnaBbaYDQBj0-oao709scndaRZifLRuzJRgNVcSzd8XsuGJ-4uBbEUbdAlWFQz_6HkFw-AT6QvtHY4sIV4zpurtJ13FIBZJw5DYnKUc0lTPGLItY507UIdcj7nBxJ2WZ7zMMYQN6FvTlEsX9vRQQ51tpjD9-9ylRac-P1jqyPFi7WtKf__I',
                'https://lh3.googleusercontent.com/aida-public/AB6AXuDAkokFDlJJU2ut7ADXOxXgokfLf-nFbve3PmdniZPK8ZjoxO1reJWexxqbxEwDy4thsHQWxI68ySeYsam7yLDI2EWIKoa7no1yZya_-OvKHcnsOU5ukziXluccC3dUHzW8BMFTNlh6bLkgI8bMSDhUTMaTR2jWDsFNRzAKs9RzZ11om4Ki5c9Xj92fIm685lMzUrq9Bk1QuG6y3TXwjzpc6HZVTEFv-aZfUUXt5s020g03CNeOFTD-_F8HtoWP1VEGiCAo5k4gC9o',
                'https://lh3.googleusercontent.com/aida-public/AB6AXuBDI76aOeylc3VRSRNZ04zYnjdju83SPGLpYh8d2uB5bogOv5w6wJcvfx6AJMDR5RQKBUrAlNQptRHNf3XJ61WGVmFOAc3-gFksvrlt6w6F_TaBfQ6Mij2zaM_KjWfcqlVkL82KYbgijsqdy-ptNMAsnQRBVHpqk2dGVjd6UsWwcuOnY9ToQPlAZwJXaT0sdGvDFkoLuAn_IIt1cP9aJJ9LXF0dXPKQK0bAJi57JnH1pgIKL4F5lFuQICzWqRy0VNT1R6BfwuOFvKg',
                'https://lh3.googleusercontent.com/aida-public/AB6AXuBHACwqs1DoKdU4I8vIetAZ1Ncf0TX2gOpmC_WxRtq2BmFEm-uJnLxeY3lrs_zfMradBduqlzQqbwRp5skRm1dQ4Ff85ob-zSZNkEhwrkNdJFgidiWuFw1ZdgMsuWEo0w82df7KisJHze0g2woOiTDZeoS_BJMCUr_VeiFECrJnlibmkGXTF67XclMyzRwRH0OjcUgdGN1GctXt_aRiiUbW4CAPTWV8RwnBLbb_HgnjSqYmbKHcclCaQQmdEbphHpEeCMk8w9JtxY4'
            ), 1, 100
        )) ELSE '' END
    ), 1, 400
);

-- 更新评价表的images字段（随机选择1-2个图片链接，用逗号分隔，限制总长度）
UPDATE game_partner_reviews SET images = SUBSTRING(
    CONCAT(
        SUBSTRING(
            ELT(
                FLOOR(RAND() * 6) + 1,
                'https://lh3.googleusercontent.com/aida-public/AB6AXuBoBTc_5bZuXQ8l_u2UKFazkITvVt5UY-tB83GE9qFMKnbb7Gz7DBuHH11MCcfExFpNociu2AurEP9Lt2NRc9nvSntdZ9hgcWNL_d-0yyyC7bLbO0F8qFUi1FZ_0xgHBG5ZWEfyBs3f5BMl_rBN4SHJoDd3xp76P8kx7eQBwXzcI46GuMySscFwGrnXs_YK9_ArHQEUVcsZUe0o_yRl84Nf4j3WwXor_Xd2gFFDgNuPdbSQuyQiPQkAovGTm7Cek_vM2ZGapACwBM4',
                'https://lh3.googleusercontent.com/aida-public/AB6AXuChoKqMjV8Zopn6t52b1X_D1IHpomYuX0uMCtUOdsEjdUbzmVSeBr-5ZKxEhOUlx1mPSxANNKH4sghnW_76nSPFe5RQDcs0D0146L6oIlgGmuRs6DcX59vrx67krFal1DduP8vpSj5o2giMgdNHYP2pIzqcKEuJfnSxyvoN_fgqJahQCNwzPm37cGoeZUvBWw_zifeh-pvUxLyRR7_8puuGkKr-5w05-1orNhrAYXS-S8PDY6s1vdzatJ05jaJSvNTo7dS88EwkP2M',
                'https://lh3.googleusercontent.com/aida-public/AB6AXuAPWd42wRhAneFPqDBu00-X6zV_3VmT3LwRUFe1rO4eEyOa-IDEPhh1HH7x2WWdyarjgr_4WzatpsdpqaJXJlMRwo9dtnaBbaYDQBj0-oao709scndaRZifLRuzJRgNVcSzd8XsuGJ-4uBbEUbdAlWFQz_6HkFw-AT6QvtHY4sIV4zpurtJ13FIBZJw5DYnKUc0lTPGLItY507UIdcj7nBxJ2WZ7zMMYQN6FvTlEsX9vRQQ51tpjD9-9ylRac-P1jqyPFi7WtKf__I',
                'https://lh3.googleusercontent.com/aida-public/AB6AXuDAkokFDlJJU2ut7ADXOxXgokfLf-nFbve3PmdniZPK8ZjoxO1reJWexxqbxEwDy4thsHQWxI68ySeYsam7yLDI2EWIKoa7no1yZya_-OvKHcnsOU5ukziXluccC3dUHzW8BMFTNlh6bLkgI8bMSDhUTMaTR2jWDsFNRzAKs9RzZ11om4Ki5c9Xj92fIm685lMzUrq9Bk1QuG6y3TXwjzpc6HZVTEFv-aZfUUXt5s020g03CNeOFTD-_F8HtoWP1VEGiCAo5k4gC9o',
                'https://lh3.googleusercontent.com/aida-public/AB6AXuBDI76aOeylc3VRSRNZ04zYnjdju83SPGLpYh8d2uB5bogOv5w6wJcvfx6AJMDR5RQKBUrAlNQptRHNf3XJ61WGVmFOAc3-gFksvrlt6w6F_TaBfQ6Mij2zaM_KjWfcqlVkL82KYbgijsqdy-ptNMAsnQRBVHpqk2dGVjd6UsWwcuOnY9ToQPlAZwJXaT0sdGvDFkoLuAn_IIt1cP9aJJ9LXF0dXPKQK0bAJi57JnH1pgIKL4F5lFuQICzWqRy0VNT1R6BfwuOFvKg',
                'https://lh3.googleusercontent.com/aida-public/AB6AXuBHACwqs1DoKdU4I8vIetAZ1Ncf0TX2gOpmC_WxRtq2BmFEm-uJnLxeY3lrs_zfMradBduqlzQqbwRp5skRm1dQ4Ff85ob-zSZNkEhwrkNdJFgidiWuFw1ZdgMsuWEo0w82df7KisJHze0g2woOiTDZeoS_BJMCUr_VeiFECrJnlibmkGXTF67XclMyzRwRH0OjcUgdGN1GctXt_aRiiUbW4CAPTWV8RwnBLbb_HgnjSqYmbKHcclCaQQmdEbphHpEeCMk8w9JtxY4'
            ), 1, 100
        ),
        CASE WHEN RAND() > 0.5 THEN CONCAT(',', SUBSTRING(
            ELT(
                FLOOR(RAND() * 6) + 1,
                'https://lh3.googleusercontent.com/aida-public/AB6AXuBoBTc_5bZuXQ8l_u2UKFazkITvVt5UY-tB83GE9qFMKnbb7Gz7DBuHH11MCcfExFpNociu2AurEP9Lt2NRc9nvSntdZ9hgcWNL_d-0yyyC7bLbO0F8qFUi1FZ_0xgHBG5ZWEfyBs3f5BMl_rBN4SHJoDd3xp76P8kx7eQBwXzcI46GuMySscFwGrnXs_YK9_ArHQEUVcsZUe0o_yRl84Nf4j3WwXor_Xd2gFFDgNuPdbSQuyQiPQkAovGTm7Cek_vM2ZGapACwBM4',
                'https://lh3.googleusercontent.com/aida-public/AB6AXuChoKqMjV8Zopn6t52b1X_D1IHpomYuX0uMCtUOdsEjdUbzmVSeBr-5ZKxEhOUlx1mPSxANNKH4sghnW_76nSPFe5RQDcs0D0146L6oIlgGmuRs6DcX59vrx67krFal1DduP8vpSj5o2giMgdNHYP2pIzqcKEuJfnSxyvoN_fgqJahQCNwzPm37cGoeZUvBWw_zifeh-pvUxLyRR7_8puuGkKr-5w05-1orNhrAYXS-S8PDY6s1vdzatJ05jaJSvNTo7dS88EwkP2M',
                'https://lh3.googleusercontent.com/aida-public/AB6AXuAPWd42wRhAneFPqDBu00-X6zV_3VmT3LwRUFe1rO4eEyOa-IDEPhh1HH7x2WWdyarjgr_4WzatpsdpqaJXJlMRwo9dtnaBbaYDQBj0-oao709scndaRZifLRuzJRgNVcSzd8XsuGJ-4uBbEUbdAlWFQz_6HkFw-AT6QvtHY4sIV4zpurtJ13FIBZJw5DYnKUc0lTPGLItY507UIdcj7nBxJ2WZ7zMMYQN6FvTlEsX9vRQQ51tpjD9-9ylRac-P1jqyPFi7WtKf__I',
                'https://lh3.googleusercontent.com/aida-public/AB6AXuDAkokFDlJJU2ut7ADXOxXgokfLf-nFbve3PmdniZPK8ZjoxO1reJWexxqbxEwDy4thsHQWxI68ySeYsam7yLDI2EWIKoa7no1yZya_-OvKHcnsOU5ukziXluccC3dUHzW8BMFTNlh6bLkgI8bMSDhUTMaTR2jWDsFNRzAKs9RzZ11om4Ki5c9Xj92fIm685lMzUrq9Bk1QuG6y3TXwjzpc6HZVTEFv-aZfUUXt5s020g03CNeOFTD-_F8HtoWP1VEGiCAo5k4gC9o',
                'https://lh3.googleusercontent.com/aida-public/AB6AXuBDI76aOeylc3VRSRNZ04zYnjdju83SPGLpYh8d2uB5bogOv5w6wJcvfx6AJMDR5RQKBUrAlNQptRHNf3XJ61WGVmFOAc3-gFksvrlt6w6F_TaBfQ6Mij2zaM_KjWfcqlVkL82KYbgijsqdy-ptNMAsnQRBVHpqk2dGVjd6UsWwcuOnY9ToQPlAZwJXaT0sdGvDFkoLuAn_IIt1cP9aJJ9LXF0dXPKQK0bAJi57JnH1pgIKL4F5lFuQICzWqRy0VNT1R6BfwuOFvKg',
                'https://lh3.googleusercontent.com/aida-public/AB6AXuBHACwqs1DoKdU4I8vIetAZ1Ncf0TX2gOpmC_WxRtq2BmFEm-uJnLxeY3lrs_zfMradBduqlzQqbwRp5skRm1dQ4Ff85ob-zSZNkEhwrkNdJFgidiWuFw1ZdgMsuWEo0w82df7KisJHze0g2woOiTDZeoS_BJMCUr_VeiFECrJnlibmkGXTF67XclMyzRwRH0OjcUgdGN1GctXt_aRiiUbW4CAPTWV8RwnBLbb_HgnjSqYmbKHcclCaQQmdEbphHpEeCMk8w9JtxY4'
            ), 1, 100
        )) ELSE '' END
    ), 1, 250
);

-- 更新游戏表的icon和image字段（使用较短的链接）
UPDATE game_partner_games SET 
    icon = SUBSTRING(
        ELT(
            FLOOR(RAND() * 6) + 1,
            'https://lh3.googleusercontent.com/aida-public/AB6AXuBoBTc_5bZuXQ8l_u2UKFazkITvVt5UY-tB83GE9qFMKnbb7Gz7DBuHH11MCcfExFpNociu2AurEP9Lt2NRc9nvSntdZ9hgcWNL_d-0yyyC7bLbO0F8qFUi1FZ_0xgHBG5ZWEfyBs3f5BMl_rBN4SHJoDd3xp76P8kx7eQBwXzcI46GuMySscFwGrnXs_YK9_ArHQEUVcsZUe0o_yRl84Nf4j3WwXor_Xd2gFFDgNuPdbSQuyQiPQkAovGTm7Cek_vM2ZGapACwBM4',
            'https://lh3.googleusercontent.com/aida-public/AB6AXuChoKqMjV8Zopn6t52b1X_D1IHpomYuX0uMCtUOdsEjdUbzmVSeBr-5ZKxEhOUlx1mPSxANNKH4sghnW_76nSPFe5RQDcs0D0146L6oIlgGmuRs6DcX59vrx67krFal1DduP8vpSj5o2giMgdNHYP2pIzqcKEuJfnSxyvoN_fgqJahQCNwzPm37cGoeZUvBWw_zifeh-pvUxLyRR7_8puuGkKr-5w05-1orNhrAYXS-S8PDY6s1vdzatJ05jaJSvNTo7dS88EwkP2M',
            'https://lh3.googleusercontent.com/aida-public/AB6AXuAPWd42wRhAneFPqDBu00-X6zV_3VmT3LwRUFe1rO4eEyOa-IDEPhh1HH7x2WWdyarjgr_4WzatpsdpqaJXJlMRwo9dtnaBbaYDQBj0-oao709scndaRZifLRuzJRgNVcSzd8XsuGJ-4uBbEUbdAlWFQz_6HkFw-AT6QvtHY4sIV4zpurtJ13FIBZJw5DYnKUc0lTPGLItY507UIdcj7nBxJ2WZ7zMMYQN6FvTlEsX9vRQQ51tpjD9-9ylRac-P1jqyPFi7WtKf__I',
            'https://lh3.googleusercontent.com/aida-public/AB6AXuDAkokFDlJJU2ut7ADXOxXgokfLf-nFbve3PmdniZPK8ZjoxO1reJWexxqbxEwDy4thsHQWxI68ySeYsam7yLDI2EWIKoa7no1yZya_-OvKHcnsOU5ukziXluccC3dUHzW8BMFTNlh6bLkgI8bMSDhUTMaTR2jWDsFNRzAKs9RzZ11om4Ki5c9Xj92fIm685lMzUrq9Bk1QuG6y3TXwjzpc6HZVTEFv-aZfUUXt5s020g03CNeOFTD-_F8HtoWP1VEGiCAo5k4gC9o',
            'https://lh3.googleusercontent.com/aida-public/AB6AXuBDI76aOeylc3VRSRNZ04zYnjdju83SPGLpYh8d2uB5bogOv5w6wJcvfx6AJMDR5RQKBUrAlNQptRHNf3XJ61WGVmFOAc3-gFksvrlt6w6F_TaBfQ6Mij2zaM_KjWfcqlVkL82KYbgijsqdy-ptNMAsnQRBVHpqk2dGVjd6UsWwcuOnY9ToQPlAZwJXaT0sdGvDFkoLuAn_IIt1cP9aJJ9LXF0dXPKQK0bAJi57JnH1pgIKL4F5lFuQICzWqRy0VNT1R6BfwuOFvKg',
            'https://lh3.googleusercontent.com/aida-public/AB6AXuBHACwqs1DoKdU4I8vIetAZ1Ncf0TX2gOpmC_WxRtq2BmFEm-uJnLxeY3lrs_zfMradBduqlzQqbwRp5skRm1dQ4Ff85ob-zSZNkEhwrkNdJFgidiWuFw1ZdgMsuWEo0w82df7KisJHze0g2woOiTDZeoS_BJMCUr_VeiFECrJnlibmkGXTF67XclMyzRwRH0OjcUgdGN1GctXt_aRiiUbW4CAPTWV8RwnBLbb_HgnjSqYmbKHcclCaQQmdEbphHpEeCMk8w9JtxY4'
        ), 1, 100
    ),
    image = SUBSTRING(
        ELT(
            FLOOR(RAND() * 6) + 1,
            'https://lh3.googleusercontent.com/aida-public/AB6AXuBoBTc_5bZuXQ8l_u2UKFazkITvVt5UY-tB83GE9qFMKnbb7Gz7DBuHH11MCcfExFpNociu2AurEP9Lt2NRc9nvSntdZ9hgcWNL_d-0yyyC7bLbO0F8qFUi1FZ_0xgHBG5ZWEfyBs3f5BMl_rBN4SHJoDd3xp76P8kx7eQBwXzcI46GuMySscFwGrnXs_YK9_ArHQEUVcsZUe0o_yRl84Nf4j3WwXor_Xd2gFFDgNuPdbSQuyQiPQkAovGTm7Cek_vM2ZGapACwBM4',
            'https://lh3.googleusercontent.com/aida-public/AB6AXuChoKqMjV8Zopn6t52b1X_D1IHpomYuX0uMCtUOdsEjdUbzmVSeBr-5ZKxEhOUlx1mPSxANNKH4sghnW_76nSPFe5RQDcs0D0146L6oIlgGmuRs6DcX59vrx67krFal1DduP8vpSj5o2giMgdNHYP2pIzqcKEuJfnSxyvoN_fgqJahQCNwzPm37cGoeZUvBWw_zifeh-pvUxLyRR7_8puuGkKr-5w05-1orNhrAYXS-S8PDY6s1vdzatJ05jaJSvNTo7dS88EwkP2M',
            'https://lh3.googleusercontent.com/aida-public/AB6AXuAPWd42wRhAneFPqDBu00-X6zV_3VmT3LwRUFe1rO4eEyOa-IDEPhh1HH7x2WWdyarjgr_4WzatpsdpqaJXJlMRwo9dtnaBbaYDQBj0-oao709scndaRZifLRuzJRgNVcSzd8XsuGJ-4uBbEUbdAlWFQz_6HkFw-AT6QvtHY4sIV4zpurtJ13FIBZJw5DYnKUc0lTPGLItY507UIdcj7nBxJ2WZ7zMMYQN6FvTlEsX9vRQQ51tpjD9-9ylRac-P1jqyPFi7WtKf__I',
            'https://lh3.googleusercontent.com/aida-public/AB6AXuDAkokFDlJJU2ut7ADXOxXgokfLf-nFbve3PmdniZPK8ZjoxO1reJWexxqbxEwDy4thsHQWxI68ySeYsam7yLDI2EWIKoa7no1yZya_-OvKHcnsOU5ukziXluccC3dUHzW8BMFTNlh6bLkgI8bMSDhUTMaTR2jWDsFNRzAKs9RzZ11om4Ki5c9Xj92fIm685lMzUrq9Bk1QuG6y3TXwjzpc6HZVTEFv-aZfUUXt5s020g03CNeOFTD-_F8HtoWP1VEGiCAo5k4gC9o',
            'https://lh3.googleusercontent.com/aida-public/AB6AXuBDI76aOeylc3VRSRNZ04zYnjdju83SPGLpYh8d2uB5bogOv5w6wJcvfx6AJMDR5RQKBUrAlNQptRHNf3XJ61WGVmFOAc3-gFksvrlt6w6F_TaBfQ6Mij2zaM_KjWfcqlVkL82KYbgijsqdy-ptNMAsnQRBVHpqk2dGVjd6UsWwcuOnY9ToQPlAZwJXaT0sdGvDFkoLuAn_IIt1cP9aJJ9LXF0dXPKQK0bAJi57JnH1pgIKL4F5lFuQICzWqRy0VNT1R6BfwuOFvKg',
            'https://lh3.googleusercontent.com/aida-public/AB6AXuBHACwqs1DoKdU4I8vIetAZ1Ncf0TX2gOpmC_WxRtq2BmFEm-uJnLxeY3lrs_zfMradBduqlzQqbwRp5skRm1dQ4Ff85ob-zSZNkEhwrkNdJFgidiWuFw1ZdgMsuWEo0w82df7KisJHze0g2woOiTDZeoS_BJMCUr_VeiFECrJnlibmkGXTF67XclMyzRwRH0OjcUgdGN1GctXt_aRiiUbW4CAPTWV8RwnBLbb_HgnjSqYmbKHcclCaQQmdEbphHpEeCMk8w9JtxY4'
        ), 1, 100
    );

-- 更新分类表的icon字段（使用较短的链接）
UPDATE game_partner_categories SET icon = SUBSTRING(
    ELT(
        FLOOR(RAND() * 6) + 1,
        'https://lh3.googleusercontent.com/aida-public/AB6AXuBoBTc_5bZuXQ8l_u2UKFazkITvVt5UY-tB83GE9qFMKnbb7Gz7DBuHH11MCcfExFpNociu2AurEP9Lt2NRc9nvSntdZ9hgcWNL_d-0yyyC7bLbO0F8qFUi1FZ_0xgHBG5ZWEfyBs3f5BMl_rBN4SHJoDd3xp76P8kx7eQBwXzcI46GuMySscFwGrnXs_YK9_ArHQEUVcsZUe0o_yRl84Nf4j3WwXor_Xd2gFFDgNuPdbSQuyQiPQkAovGTm7Cek_vM2ZGapACwBM4',
        'https://lh3.googleusercontent.com/aida-public/AB6AXuChoKqMjV8Zopn6t52b1X_D1IHpomYuX0uMCtUOdsEjdUbzmVSeBr-5ZKxEhOUlx1mPSxANNKH4sghnW_76nSPFe5RQDcs0D0146L6oIlgGmuRs6DcX59vrx67krFal1DduP8vpSj5o2giMgdNHYP2pIzqcKEuJfnSxyvoN_fgqJahQCNwzPm37cGoeZUvBWw_zifeh-pvUxLyRR7_8puuGkKr-5w05-1orNhrAYXS-S8PDY6s1vdzatJ05jaJSvNTo7dS88EwkP2M',
        'https://lh3.googleusercontent.com/aida-public/AB6AXuAPWd42wRhAneFPqDBu00-X6zV_3VmT3LwRUFe1rO4eEyOa-IDEPhh1HH7x2WWdyarjgr_4WzatpsdpqaJXJlMRwo9dtnaBbaYDQBj0-oao709scndaRZifLRuzJRgNVcSzd8XsuGJ-4uBbEUbdAlWFQz_6HkFw-AT6QvtHY4sIV4zpurtJ13FIBZJw5DYnKUc0lTPGLItY507UIdcj7nBxJ2WZ7zMMYQN6FvTlEsX9vRQQ51tpjD9-9ylRac-P1jqyPFi7WtKf__I',
        'https://lh3.googleusercontent.com/aida-public/AB6AXuDAkokFDlJJU2ut7ADXOxXgokfLf-nFbve3PmdniZPK8ZjoxO1reJWexxqbxEwDy4thsHQWxI68ySeYsam7yLDI2EWIKoa7no1yZya_-OvKHcnsOU5ukziXluccC3dUHzW8BMFTNlh6bLkgI8bMSDhUTMaTR2jWDsFNRzAKs9RzZ11om4Ki5c9Xj92fIm685lMzUrq9Bk1QuG6y3TXwjzpc6HZVTEFv-aZfUUXt5s020g03CNeOFTD-_F8HtoWP1VEGiCAo5k4gC9o',
        'https://lh3.googleusercontent.com/aida-public/AB6AXuBDI76aOeylc3VRSRNZ04zYnjdju83SPGLpYh8d2uB5bogOv5w6wJcvfx6AJMDR5RQKBUrAlNQptRHNf3XJ61WGVmFOAc3-gFksvrlt6w6F_TaBfQ6Mij2zaM_KjWfcqlVkL82KYbgijsqdy-ptNMAsnQRBVHpqk2dGVjd6UsWwcuOnY9ToQPlAZwJXaT0sdGvDFkoLuAn_IIt1cP9aJJ9LXF0dXPKQK0bAJi57JnH1pgIKL4F5lFuQICzWqRy0VNT1R6BfwuOFvKg',
        'https://lh3.googleusercontent.com/aida-public/AB6AXuBHACwqs1DoKdU4I8vIetAZ1Ncf0TX2gOpmC_WxRtq2BmFEm-uJnLxeY3lrs_zfMradBduqlzQqbwRp5skRm1dQ4Ff85ob-zSZNkEhwrkNdJFgidiWuFw1ZdgMsuWEo0w82df7KisJHze0g2woOiTDZeoS_BJMCUr_VeiFECrJnlibmkGXTF67XclMyzRwRH0OjcUgdGN1GctXt_aRiiUbW4CAPTWV8RwnBLbb_HgnjSqYmbKHcclCaQQmdEbphHpEeCMk8w9JtxY4'
    ), 1, 100
);

-- 查看更新结果
SELECT 'Users' as table_name, COUNT(*) as updated_rows FROM game_partner_users WHERE avatar != '';
SELECT 'Playmates' as table_name, COUNT(*) as updated_rows FROM game_partner_playmates WHERE avatar != '';
SELECT 'Community Posts' as table_name, COUNT(*) as updated_rows FROM game_partner_community_posts WHERE images != '';
SELECT 'Reviews' as table_name, COUNT(*) as updated_rows FROM game_partner_reviews WHERE images != '';
SELECT 'Games' as table_name, COUNT(*) as updated_rows FROM game_partner_games WHERE icon != '' AND image != '';
SELECT 'Categories' as table_name, COUNT(*) as updated_rows FROM game_partner_categories WHERE icon != '';