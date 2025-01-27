-- Copyright 2025 Specter Ops, Inc.
--
-- Licensed under the Apache License, Version 2.0
-- you may not use this file except in compliance with the License.
-- You may obtain a copy of the License at
--
--     http://www.apache.org/licenses/LICENSE-2.0
--
-- Unless required by applicable law or agreed to in writing, software
-- distributed under the License is distributed on an "AS IS" BASIS,
-- WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
-- See the License for the specific language governing permissions and
-- limitations under the License.
--
-- SPDX-License-Identifier: Apache-2.0

-- case: with 1 as target match (n:NodeKind1) where n.value = target return n
with s0 as (select 1 as target)
with s1 as (select s0.target as target, (n0.id, n0.kind_ids, n0.properties)::nodecomposite as n0
            from s0,
                 node n0
            where n0.kind_ids operator (pg_catalog.&&) array [1]::int2[]
              and n0.properties ->> 'value' = target)
select s1.n0 as n
from s1;

-- exclusive:
-- case: match (n:NodeKind1) where n.value = 1 with n match (b) where id(b) = id(n) return b
;
