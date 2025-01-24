// Copyright 2023 Specter Ops, Inc.
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

import { Box, CircularProgress } from '@mui/material';
import makeStyles from '@mui/styles/makeStyles';
import { GenericErrorBoundaryFallback, apiClient } from 'bh-shared-ui';
import React, { Suspense, useEffect } from 'react';
import { ErrorBoundary } from 'react-error-boundary';
import { Route, Routes } from 'react-router-dom';
import AuthenticatedRoute from 'src/components/AuthenticatedRoute';
import { ListAssetGroups } from 'src/ducks/assetgroups/actionCreators';
import { fullyAuthenticatedSelector } from 'src/ducks/auth/authSlice';
import { fetchAssetGroups, setDomain } from 'src/ducks/global/actions';
import { ROUTES } from 'src/routes';
import { useAppDispatch, useAppSelector } from 'src/store';

const useStyles = makeStyles({
    content: {
        position: 'relative',
        width: '100%',
        height: '100%',
        minHeight: '100%',
    },
});

const Content: React.FC = () => {
    const classes = useStyles();
    const dispatch = useAppDispatch();
    const authState = useAppSelector((state) => state.auth);
    const isFullyAuthenticated = useAppSelector(fullyAuthenticatedSelector);

    useEffect(() => {
        if (isFullyAuthenticated) {
            dispatch(fetchAssetGroups());
            dispatch(ListAssetGroups());
        }
    }, [authState, isFullyAuthenticated, dispatch]);

    // set inital domain/tenant once user is authenticated
    useEffect(() => {
        if (isFullyAuthenticated) {
            const ctrl = new AbortController();
            apiClient
                .getAvailableDomains({ signal: ctrl.signal })
                .then((result) => {
                    const collectedDomains = result.data.data
                        // omit uncollected domains
                        .filter((domain: any) => domain.collected)
                        // sort by impactValue descending
                        .sort((a: any, b: any) => b.impactValue - a.impactValue);
                    if (collectedDomains.length > 0) {
                        dispatch(setDomain(collectedDomains[0]));
                    } else {
                        dispatch(setDomain(null));
                    }
                })
                .catch(() => {
                    dispatch(setDomain(null));
                });
            return () => ctrl.abort();
        }
    }, [isFullyAuthenticated, dispatch]);

    return (
        <Box className={classes.content}>
            <Suspense
                fallback={
                    <Box
                        position='absolute'
                        top='0'
                        left='0'
                        right='0'
                        bottom='0'
                        display='flex'
                        alignItems='center'
                        justifyContent='center'
                        zIndex={1000}>
                        <CircularProgress color='primary' size={80} />
                    </Box>
                }>
                <Routes>
                    {ROUTES.map((route) => {
                        return route.authenticationRequired ? (
                            <Route
                                path={route.path}
                                element={
                                    // To do: Tokenize the gutter value, because its the amount the admin menu needs to move left so something like pl-nav
                                    // Note: We add a left padding value to account for pages that have nav bar, h-full is because when adding the div it collapsed the views
                                    <ErrorBoundary fallbackRender={GenericErrorBoundaryFallback}>
                                        <AuthenticatedRoute>
                                            <div className={`h-full ${route.navigation && 'pl-[56px]'} `}>
                                                <route.component />
                                            </div>
                                        </AuthenticatedRoute>
                                    </ErrorBoundary>
                                }
                                key={route.path}
                            />
                        ) : (
                            <Route
                                path={route.path}
                                element={
                                    <ErrorBoundary fallbackRender={GenericErrorBoundaryFallback}>
                                        <route.component />
                                    </ErrorBoundary>
                                }
                                key={route.path}
                            />
                        );
                    })}
                </Routes>
            </Suspense>
        </Box>
    );
};

export default Content;
