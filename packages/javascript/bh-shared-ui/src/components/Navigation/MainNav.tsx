// Copyright 2024 Specter Ops, Inc.
//
// Licensed under the Apache License, Version 2.0
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// SPDX-License-Identifier: Apache-2.0

import { FC, useState, ReactNode } from 'react';
import { useLocation, Link as RouterLink } from 'react-router-dom';
import { useApiVersion } from '../../hooks';

const MainNavLogoTextImage: FC<{ mainNavLogoData: any }> = ({ mainNavLogoData }) => {
    return (
        <img
            src={mainNavLogoData.image.imageUrl}
            alt={mainNavLogoData.image.altText}
            height={mainNavLogoData.image.dimensions.height}
            width={mainNavLogoData.image.dimensions.width}
            className={mainNavLogoData.image.classes}
        />
    );
};

const MainNavListItem: FC<{ children: ReactNode; route?: string }> = ({ children, route }) => {
    const location = useLocation();
    const isActiveRoute = route ? location.pathname.includes(route.replace(/\*/g, '')) : false;

    return (
        <li
            className={`h-10 px-2 mx-2 flex items-center ${isActiveRoute ? 'text-primary bg-neutral-light-4' : 'hover:text-secondary hover:underline'} cursor-pointer rounded`}>
            {children}
        </li>
    );
};

const MainNavItemAction: FC<{ onClick: () => void; children: ReactNode; isMenuExpanded: boolean }> = ({
    onClick,
    children,
    isMenuExpanded,
}) => {
    return (
        // Note: The w-full terniary is to avoid the hover area to overflow out of the nav when its collapsed
        <button
            onClick={onClick}
            className={`h-10 ${isMenuExpanded ? 'w-full' : 'w-auto'} absolute left-5 flex items-center gap-x-2`}>
            {children}
        </button>
    );
};

const MainNavItemLink: FC<{ route: string; children: ReactNode; isMenuExpanded: boolean }> = ({
    route,
    children,
    isMenuExpanded,
}) => {
    return (
        // Note: The w-full terniary is to avoid the hover area to overflow out of the nav when its collapsed
        <RouterLink
            to={route as string}
            className={`h-10 ${isMenuExpanded ? 'w-full' : 'w-auto'} absolute left-5 flex items-center gap-x-2`}>
            {children}
        </RouterLink>
    );
};

const MainNavItemLabel: FC<{ icon: ReactNode; label: ReactNode | string; isMenuExpanded: boolean }> = ({
    icon,
    label,
    isMenuExpanded,
}) => {
    return (
        // Note: The min-h here is to keep spacing between the logo and the list below.
        <>
            <span className='flex'>{icon}</span>
            <span
                className={`whitespace-nowrap flex min-h-10 items-center gap-x-5 font-medium text-xl ${isMenuExpanded ? 'opacity-100 block' : 'opacity-0 hidden'} duration-200 ease-in`}>
                {label}
            </span>
        </>
    );
};

const MainNavVersionNumber: FC<{ isMenuExpanded: boolean }> = ({ isMenuExpanded }) => {
    const { data: apiVersionResponse, isSuccess } = useApiVersion();
    const apiVersion = isSuccess && apiVersionResponse?.server_version;

    return (
        // Note: The min-h allows for the version number to keep its position when the nav is scrollable
        <div className='relative w-full flex min-h-10 h-10 overflow-x-hidden'>
            <div
                className={`w-full flex absolute bottom-3 ${isMenuExpanded ? 'left-16' : 'left-4'} duration-300 ease-in-out text-xs whitespace-nowrap font-medium text-neutral-dark-0 dark:text-neutral-light-1`}>
                <span
                    className={`${isMenuExpanded ? 'opacity-100 block' : 'opacity-0 hidden'} duration-300 ease-in-out`}>
                    BloodHound:&nbsp;
                </span>
                <span className={`${!isMenuExpanded && 'max-w-9 overflow-x-hidden'}`}>{apiVersion}</span>
            </div>
        </div>
    );
};

const MainNav: FC<{ mainNavData: any }> = ({ mainNavData }) => {
    const [isMenuExpanded, setIsMenuExpanded] = useState(false);

    return (
        // Note: z-index needs to be higher than sub-nav
        <nav
            className={`z-[1201] fixed top-0 left-0 h-full ${isMenuExpanded ? 'w-72 overflow-y-auto overflow-x-hidden' : 'w-16'} duration-300 ease-in flex flex-col items-center pt-4  bg-neutral-light-2 text-neutral-dark-0 dark:bg-neutral-dark-2 dark:text-neutral-light-1 print:hidden shadow-sm`}
            onMouseEnter={() => setIsMenuExpanded(true)}
            onMouseLeave={() => setIsMenuExpanded(false)}>
            <MainNavItemLink route={mainNavData.logo.route} isMenuExpanded={isMenuExpanded}>
                <MainNavItemLabel
                    icon={mainNavData.logo.icon}
                    label={<MainNavLogoTextImage mainNavLogoData={mainNavData.logo} />}
                    isMenuExpanded={isMenuExpanded}
                />
            </MainNavItemLink>
            {/* Note: min height here is to keep the version number in bottom of nav */}
            <div className='h-full min-h-[700px] w-full flex flex-col justify-between mt-6'>
                <ul className='flex flex-col gap-6 mt-8'>
                    {mainNavData.primaryList.map((listDataItem: any) => (
                        <MainNavListItem key={listDataItem.testid} route={listDataItem.route as string}>
                            <MainNavItemLink route={listDataItem.route as string} isMenuExpanded={isMenuExpanded}>
                                <MainNavItemLabel
                                    icon={listDataItem.icon}
                                    label={listDataItem.label}
                                    isMenuExpanded={isMenuExpanded}
                                />
                            </MainNavItemLink>
                        </MainNavListItem>
                    ))}
                </ul>
                <ul className='flex flex-col gap-6 mt-16'>
                    {mainNavData.secondaryList.map((listDataItem: any) =>
                        listDataItem.route ? (
                            <MainNavListItem key={listDataItem.testid} route={listDataItem.route as string}>
                                <MainNavItemLink route={listDataItem.route as string} isMenuExpanded={isMenuExpanded}>
                                    <MainNavItemLabel
                                        icon={listDataItem.icon}
                                        label={listDataItem.label}
                                        isMenuExpanded={isMenuExpanded}
                                    />
                                </MainNavItemLink>
                            </MainNavListItem>
                        ) : (
                            <MainNavListItem key={listDataItem.testid}>
                                <MainNavItemAction
                                    onClick={(() => listDataItem.functionHandler as () => void)()}
                                    isMenuExpanded={isMenuExpanded}>
                                    <MainNavItemLabel
                                        icon={listDataItem.icon}
                                        label={listDataItem.label}
                                        isMenuExpanded={isMenuExpanded}
                                    />
                                </MainNavItemAction>
                            </MainNavListItem>
                        )
                    )}
                </ul>
            </div>
            <MainNavVersionNumber isMenuExpanded={isMenuExpanded} />
        </nav>
    );
};

export default MainNav;
